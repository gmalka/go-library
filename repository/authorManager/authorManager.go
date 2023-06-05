package authorManager

import (
	"go-library/model"

	"github.com/jmoiron/sqlx"
)

type authorManager struct {
	db *sqlx.DB
}

type AuthorManagerI interface {
	Add(author model.Author) error
	Get(id int) (model.Author, error)
	GetAll() ([]model.Author, error)
	Delete(id int) error
}

func NewAuthorManager(db *sqlx.DB) AuthorManagerI {
	return authorManager{db: db}
}

func (a authorManager) Delete(id int) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM author WHERE id=$1", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM books WHERE author_id=$1", id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a authorManager) Get(id int) (model.Author, error) {
	row := a.db.QueryRow("SELECT * FROM author WHERE id=$1", id)

	author := model.Author{}
	err := row.Scan(&author.Id, &author.Name)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (a authorManager) GetAll() ([]model.Author, error) {
	rows, err := a.db.Query("SELECT * FROM author")
	if err != nil {
		return nil, err
	}

	authors := make([]model.Author, 0, 10)
	for rows.Next() {
		author := model.Author{}

		err = rows.Scan(&author.Id, &author.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a authorManager) Add(author model.Author) error {
	_, err := a.db.Exec("INSERT INTO author(name) VALUES($1)", author.Name)
	if err != nil {
		return err
	}

	return nil
}