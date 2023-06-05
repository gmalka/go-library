package main

import (
	"context"
	"fmt"
	"go-library/handler"
	"go-library/repository/authorManager"
	"go-library/repository/bookManager"
	"go-library/repository/libraryManager"
	"go-library/repository/userManager"
	"go-library/service/libraryService"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresUrl := os.Getenv("POSTGRES_URL")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB")
	sslmode := os.Getenv("POSTGRES_SSL")

	url := os.Getenv("URL_TO_RUN")
	port := os.Getenv("PORT_TO_RUN")

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		postgresUrl, postgresPort, postgresUser, databaseName, postgresPassword, sslmode))
	if err != nil {
		log.Fatalln(err)
	}
	err = generateData(db)
	if err != nil {
		log.Fatalln(err)
	}

	authManager := authorManager.NewAuthorManager(db)
	userManager := userManager.NewuserManager(db)
	bookManager := bookManager.NewBookManager(db)
	libraryManager := libraryManager.NewLibraryManager(db)

	service := libraryService.NewLibrary(authManager, userManager, bookManager, libraryManager)

	
	h := handler.NewHandler(service)
	RunServer(fmt.Sprintf("%s:%s", url, port), h.IniRouter())
}

func generateData(db *sqlx.DB) error {
	res, err := db.Query("SELECT * FROM author")
	if err != nil {
		return err
	}
	if !res.Next() {
		for i := 0; i < 10; i++ {
			name := gofakeit.Name()
			lastname := gofakeit.LastName()
			db.Exec("INSERT INTO author(name) VALUES($1)", fmt.Sprintf("%s %s", name, lastname))
		}
	}

	res, err = db.Query("SELECT * FROM books")
	if err != nil {
		return err
	}
	if !res.Next() {
		for i := 0; i < 100; i++ {
			auth_id := 0
			row := db.QueryRow("SELECT id FROM author ORDER BY RANDOM() LIMIT 1")
			err = row.Scan(&auth_id)
			if err != nil {
				return err
			}
			name1 := gofakeit.Word()
			name2 := gofakeit.Word()

			_, err = db.Exec("INSERT INTO books(name, author_id) VALUES($1, $2)", fmt.Sprintf("%s %s", name1, name2), auth_id)
			if err != nil {
				return err
			}
		}
	}

	res, err = db.Query("SELECT * FROM users")
	if err != nil {
		return err
	}
	if !res.Next() {
		for i := 0; i < 50; i++ {
			name := gofakeit.Username()

			db.Exec("INSERT INTO users(name) VALUES($1)", name)
		}
	}

	return nil
}

func RunServer(addr string, h http.Handler) {
	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	go srv.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server...")

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	srv.Shutdown(ctx)
	log.Println("Server exited")
}


// ---------------------->> MIGRATE ---------------------->>
// migrate -path db/migrate/ -database "postgres://gmalka:postgres@localhost:5432/mytestdb?sslmode=disable" up

// ---------------------->> DATABASE --------------------->>
// docker run --rm -d --name postgres -p 5432:5432 -e POSTGRES_USER=gmalka -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=mytestdb postgres