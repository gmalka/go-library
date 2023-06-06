package model

type Author struct {
	Id int				`json:"id,omitempty"`
	Name string			`json:"name"`
}

type AuthorWithBooks struct {
	Id int				`json:"id,omitempty"`
	Name string			`json:"name"`
	Books []Book		`json:"books"`
}