package libraryService

import (
	"go-library/model"
)

type AuthorManagerI interface {
	Add(author model.Author) error
	Get(id int) (model.Author, error)
	GetAll() ([]model.Author, error)
	Delete(id int) error
}

type UserManagerI interface {
	Add(user model.User) error
	Get(id int) (model.User, error)
	GetAll() ([]model.User, error)
	Delete(id int) error
}

type BookManagerI interface {
	Add(book model.Book) error
	Get(id int) (model.Book, error)
	GetAll() ([]model.Book, error)
	Delete(id int) error
}

type LibraryMangerI interface {
	PutBookBack(userId, bookId int) error
	TakeBook(userId, bookId int) error
	GetAllTakenBooks() ([]model.TakenBook, error)
	GetAllTakenBooksOfUser(id int) ([]model.Book, error)
}

type Library interface {
	PutBookBack(userId, bookId int) error
	TakeBook(userId, bookId int) error
	GetAllTakenBooks() ([]model.TakenBook, error)
	GetAllTakenBooksOfUser(id int) ([]model.Book, error)

	AddUser(user model.User) error
	GetUser(id int) (model.User, error)
	GetAllUsers() ([]model.User, error)
	DeleteUser(id int) error

	AddAuthor(author model.Author) error
	GetAuthor(id int) (model.Author, error)
	GetAllAuthors() ([]model.Author, error)
	DeleteAuthor(id int) error

	AddBook(book model.Book) error
	GetBook(id int) (model.Book, error)
	GetAllBooks() ([]model.Book, error)
	DeleteBook(id int) error
}

type myLibrary struct {
	authorManager  AuthorManagerI
	booksManager   BookManagerI
	usersManager   UserManagerI
	libraryManager LibraryMangerI
}

func NewLibrary(a AuthorManagerI, u UserManagerI, b BookManagerI, l LibraryMangerI) Library {
	return myLibrary{authorManager: a, usersManager: u, booksManager: b, libraryManager: l}
}

// <-------------LIBRARY------------->

func (m myLibrary) GetAllTakenBooksOfUser(id int) ([]model.Book, error) {
	return m.libraryManager.GetAllTakenBooksOfUser(id)
}

func (m myLibrary) PutBookBack(userId, bookId int) error {
	return m.libraryManager.PutBookBack(userId, bookId)
}

func (m myLibrary) TakeBook(userId, bookId int) error {
	return m.libraryManager.TakeBook(userId, bookId)
}

func (m myLibrary) GetAllTakenBooks() ([]model.TakenBook, error) {
	return m.libraryManager.GetAllTakenBooks()
}

// <-------------USER------------->

func (m myLibrary) AddUser(user model.User) error {
	return m.usersManager.Add(user)
}

func (m myLibrary) GetUser(id int) (model.User, error) {
	return m.usersManager.Get(id)
}

func (m myLibrary) GetAllUsers() ([]model.User, error) {
	return m.usersManager.GetAll()
}

func (m myLibrary) DeleteUser(id int) error {
	return m.usersManager.Delete(id)
}

// <-------------AUTHOR------------->

func (m myLibrary) AddAuthor(author model.Author) error {
	return m.authorManager.Add(author)
}

func (m myLibrary) GetAuthor(id int) (model.Author, error) {
	return m.authorManager.Get(id)
}

func (m myLibrary) GetAllAuthors() ([]model.Author, error) {
	return m.authorManager.GetAll()
}

func (m myLibrary) DeleteAuthor(id int) error {
	return m.authorManager.Delete(id)
}

// <-------------BOOKS------------->

func (m myLibrary) AddBook(book model.Book) error {
	return m.booksManager.Add(book)
}

func (m myLibrary) GetBook(id int) (model.Book, error) {
	return m.booksManager.Get(id)
}

func (m myLibrary) GetAllBooks() ([]model.Book, error) {
	return m.booksManager.GetAll()
}

func (m myLibrary) DeleteBook(id int) error {
	return m.booksManager.Delete(id)
}