package userManager

import (
	"go-library/model"

	"github.com/jmoiron/sqlx"
)

type userManager struct {
	db *sqlx.DB
}

type UserManagerI interface {
	Add(user model.User) error
	Get(id int) (model.User, error)
	GetAll() ([]model.User, error)
	Delete(id int) error
}

func NewuserManager(db *sqlx.DB) UserManagerI {
	return userManager{db: db}
}

func (a userManager) Delete(id int) error {
	_, err := a.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

func (a userManager) GetAll() ([]model.User, error) {
	users := make([]model.User, 0, 10)

	rows, err := a.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (a userManager) Get(id int) (model.User, error) {
	user := model.User{}

	row := a.db.QueryRow("SELECT * FROM users WHERE id=$1", id)
	
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (a userManager) Add(user model.User) error {
	_, err := a.db.Exec("INSERT INTO users(name) VALUES($1)", user.Name)
	return err
}