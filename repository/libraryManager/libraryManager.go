package libraryManager

import (
	"database/sql"
	"errors"
	"go-library/model"

	"github.com/jmoiron/sqlx"
)

type libraryManger struct {
	db *sqlx.DB
}



type LibraryMangerI interface {
	PutBookBack(userId, bookId int) error
	TakeBook(userId, bookId int) error
	GetAllTakenBooks() ([]model.TakenBook, error)
	GetAllTakenBooksOfUser(id int) ([]model.BookWithAuthor, error)
}

func NewLibraryManager(db *sqlx.DB) LibraryMangerI {
	return libraryManger{db: db}
}

func (a libraryManger) PutBookBack(userId, bookId int) error {
	var (
		s, r interface{}
	)
	row := a.db.QueryRow("SELECT * FROM takenBook WHERE user_id=$1 AND book_id=$2", userId, bookId)
	err := row.Scan(&s, &r)
	if err != nil {
		return errors.New("can't find book")
	}

	_, err = a.db.Exec("DELETE FROM takenBook WHERE user_id=$1 AND book_id=$2", userId, bookId)
	return err
}

func (a libraryManger) TakeBook(userId, bookId int) error {
	row := a.db.QueryRow("SELECT * FROM takenBook WHERE book_id=$1", bookId)
	var s interface{}
	err := row.Scan(&s)
	if err != sql.ErrNoRows {
		return errors.New("book already was taken")
	}

	_, err = a.db.Exec("INSERT INTO takenBook(user_id, book_id) SELECT $1, $2 WHERE NOT EXISTS (SELECT 1 FROM takenBook WHERE book_id=$2)", userId, bookId)
	return err
}

func (a libraryManger) GetAllTakenBooks() ([]model.TakenBook, error) {
	rows, err := a.db.Query("SELECT * FROM takenBook")
	if err != nil {
		return nil, err
	}

	takenBooks := make([]model.TakenBook, 0, 10)
	for rows.Next() {
		var (
			user 	model.User
			book 	model.BookWithAuthor
			userId int
			bookId int
		)
		err = rows.Scan(&userId, &bookId)
		if err != nil {
			return nil, err
		}

		userRow := a.db.QueryRow("SELECT * FROM users WHERE id=$1", userId)
		
		err = userRow.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}

		bookRow := a.db.QueryRow("SELECT books.id, books.name, author.id, author.name FROM books LEFT JOIN author On author.id=books.author_id WHERE books.id=$1", bookId)
		
		err = bookRow.Scan(&book.Id, &book.Name, &book.Auth.Id, &book.Auth.Name)
		if err != nil {
			return nil, err
		}

		takenBooks = append(takenBooks, model.TakenBook{UserWhoTake: user, TakenBook: book})
	}

	return takenBooks, nil
}

func (a libraryManger) GetAllTakenBooksOfUser(id int) ([]model.BookWithAuthor, error) {
	rows, err := a.db.Query("SELECT book_id FROM takenBook WHERE user_id=$1", id)
	if err != nil {
		return nil, err
	}

	books := make([]model.BookWithAuthor, 0, 10)
	for rows.Next() {
		var (
			book 	model.BookWithAuthor
			bookId int
		)
		err = rows.Scan(&bookId)
		if err != nil {
			return nil, err
		}

		userRow := a.db.QueryRow("SELECT books.id, books.name, author.id, author.name FROM books LEFT JOIN author On author.id=books.author_id WHERE books.id=$1", bookId)
		
		err = userRow.Scan(&book.Id, &book.Name, &book.Auth.Id, &book.Auth.Name)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}