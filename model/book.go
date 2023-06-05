package model

type Book struct {
	Id int			`json:"id,omitempty"`
	Name string		`json:"name"`
	Auth Author		`json:"author"`
}