package handler

import (
	"go-library/model"
	"net/http"
	"text/template"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	swaggerTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-standalone-preset.js"></script>
    <!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-standalone-preset.js"></script> -->
    <script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
    <!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-bundle.js"></script> -->
    <link rel="stylesheet" href="//unpkg.com/swagger-ui-dist@3/swagger-ui.css" />
    <!-- <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui.css" /> -->
	<style>
		body {
			margin: 0;
		}
	</style>
    <title>Swagger</title>
</head>
<body>
    <div id="swagger-ui"></div>
    <script>
        window.onload = function() {
          SwaggerUIBundle({
            url: "/public/swagger.json?{{.Time}}",
            dom_id: '#swagger-ui',
            presets: [
              SwaggerUIBundle.presets.apis,
              SwaggerUIStandalonePreset
            ],
            layout: "StandaloneLayout"
          })
        }
    </script>
</body>
</html>
`
)


type Library interface {
	PutBookBack(userId, bookId int) error
	TakeBook(userId, bookId int) error
	GetAllTakenBooks() ([]model.TakenBook, error)
	GetAllTakenBooksOfUser(id int) ([]model.BookWithAuthor, error)

	AddUser(user model.User) error
	GetUser(id int) (model.User, error)
	GetAllUsers() ([]model.User, error)
	DeleteUser(id int) error

	AddAuthor(author model.Author) error
	GetAuthor(id int) (model.Author, error)
	GetAllAuthors() ([]model.AuthorWithBooks, error)
	DeleteAuthor(id int) error

	AddBook(book model.BookWithAuthor) error
	GetBook(bid, aid int) (model.BookWithAuthor, error)
	GetAllBooks() ([]model.BookWithAuthor, error)
	GetAllBooksOfAuthor(aid int) ([]model.Book, error)
	DeleteBook(bid, aid int) error
}

type Handler struct {
	service Library
}

func NewHandler(service Library) *Handler {
	return &Handler{service: service}
}

func (h *Handler) IniRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)


	r.Get("/swagger", h.swaggerUI)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
	})

	r.Get("/", h.GetAllTakenBooks)
	r.Get("/books", h.GetAllBooks)

	r.Route("/authors", func(r chi.Router) {
		r.Get("/", h.GetAllAuthors)
		r.Post("/", h.AddAuthor)
		r.Delete("/{id}", h.DeleteAuthor)

		r.Route("/{author_id}/books", func(r chi.Router) {
			r.Get("/", h.GetAllAuthorBooks)
			r.Post("/", h.AddBook)
			r.Delete("/{id}", h.DeleteBook)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.GetAllUsers)
		r.Post("/", h.AddUser)
		r.Delete("/{id}", h.DeleteUser)

		r.Route("/{id}/library", func(r chi.Router) {
			r.Get("/", h.GetAllTakenBooksOfUser)
			r.Post("/{book_id}", h.TakeBook)
			r.Delete("/{book_id}", h.ReturnBook)
		})
	})

	return r
}

func (h Handler)swaggerUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.New("swagger").Parse(swaggerTemplate)
	if err != nil {
		return
	}

	err = tmpl.Execute(w, struct {
		Time int64
	}{
		Time: time.Now().Unix(),
	})
	if err != nil {
		return
	}
}