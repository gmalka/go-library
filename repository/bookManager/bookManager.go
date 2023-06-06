package bookManager

import (
	"errors"
	"go-library/model"

	"github.com/jmoiron/sqlx"
)

type bookManager struct {
	db *sqlx.DB
}

type BookManagerI interface {
	Add(book model.BookWithAuthor) error
	Get(bid, aid int) (model.BookWithAuthor, error)
	GetAll() ([]model.BookWithAuthor, error)
	Delete(bid, aid int) error
	GetAllOfAuthor(aid int) ([]model.Book, error)
}

func NewBookManager(db *sqlx.DB) BookManagerI {
	return bookManager{db: db}
}

func (a bookManager) Add(book model.BookWithAuthor) error {
	_, err := a.db.Exec("INSERT INTO books(name, author_id) VALUES($1, $2)", book.Name, book.Auth.Id)
	if err != nil {
		return err
	}

	return nil
}

func (a bookManager) GetAll() ([]model.BookWithAuthor, error) {
	rows, err := a.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]model.BookWithAuthor, 0, 10)
	for rows.Next() {
		book := model.BookWithAuthor{}
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

func (a bookManager) Delete(bid, aid int) error {
	result, err := a.db.Exec("DELETE FROM books WHERE id=$1 AND author_id=$2", bid, aid)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i != 1 {
		return errors.New("cant find book")
	}

	return nil
}

func (a bookManager) Get(bid, aid int) (model.BookWithAuthor, error) {
	row := a.db.QueryRow("SELECT * FROM books WHERE id=$1 AND author_id=$2", bid, aid)

	book := model.BookWithAuthor{}
	err := row.Scan(&book.Id, &book.Name)
	if err != nil {
		return model.BookWithAuthor{}, err
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

func (a bookManager) GetAllOfAuthor(aid int) ([]model.Book, error) {
	rows, err := a.db.Query("SELECT id, name FROM books WHERE author_id=$1", aid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]model.Book, 0, 10)
	for rows.Next() {
		book := model.Book{}
		err := rows.Scan(&book.Id, &book.Name)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}