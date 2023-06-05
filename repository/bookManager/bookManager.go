package bookManager

import (
	"go-library/model"

	"github.com/jmoiron/sqlx"
)

type bookManager struct {
	db *sqlx.DB
}

type BookManagerI interface {
	Add(book model.Book) error
	Get(id int) (model.Book, error)
	GetAll() ([]model.Book, error)
	Delete(id int) error
}

func NewBookManager(db *sqlx.DB) BookManagerI {
	return bookManager{db: db}
}

func (a bookManager) Add(book model.Book) error {
	_, err := a.db.Exec("INSERT INTO books(name, author_id) VALUES($1, $2)", book.Name, book.Auth.Id)
	if err != nil {
		return err
	}

	return nil
}

func (a bookManager) GetAll() ([]model.Book, error) {
	rows, err := a.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	books := make([]model.Book, 0, 10)
	for rows.Next() {
		book := model.Book{}
		authorId := 0
		err := rows.Scan(&book.Id, &book.Name, &authorId)
		if err != nil {
			return nil, err
		}

		authrRow := a.db.QueryRow("SELECT * FROM author WHERE id=$1", authorId)
		author := model.Author{}
		err = authrRow.Scan(&author.Id, &author.Name)
		if err != nil {
			return nil, err
		}
		
		book.Auth = author

		books = append(books, book)
	}

	return books, nil
}

func (a bookManager) Delete(id int) error {
	_, err := a.db.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func (a bookManager) Get(id int) (model.Book, error) {
	row := a.db.QueryRow("SELECT * FROM books WHERE id=$1", id)

	book := model.Book{}
	err := row.Scan(&book.Id, &book.Name)
	if err != nil {
		return model.Book{}, err
	}

	authorId := 0
	err = row.Scan(&authorId)
	if err != nil {
		return book, err
	}

	authrRow := a.db.QueryRow("SELECT * FROM author WHERE id=$1", authorId)
	author := model.Author{}
	err = authrRow.Scan(&author.Id, &author.Name)
	if err != nil {
		return book, err
	}
	book.Auth = author

	return book, nil
}