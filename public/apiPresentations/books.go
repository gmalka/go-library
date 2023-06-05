package apiPresentations

import "go-library/model"

// swagger:parameters bookEmptyRequest
type bookEmptyRequest struct {}

// swagger:response bookEmptyResponse
type bookEmptyResponse struct {}


// swagger:route GET /authors/{author_id}/books books bookGetAllOfAuthorRequest
// Получение всех книг автора
// responses:
//   200: bookGetAllOfAuthorResponse

// swagger:parameters bookGetAllOfAuthorRequest
type bookGetAllOfAuthorRequest struct {
	// ID автора
	//
	// in:path
	ID_author string `json:"author_id"`
}

// swagger:response bookGetAllOfAuthorResponse
type bookGetAllOfAuthorResponse struct {
	// Массив книг
	//
	// in:body
	Body []model.Book
}


// swagger:route POST /authors/{author_id}/books books bookAddRequest
// Создание новой книги автора
// responses:
//   200: bookEmptyResponse

// swagger:parameters bookAddRequest
type bookAddRequest struct {
	// ID автора
	//
	// in:path
	ID_author string `json:"author_id"`
	// Требуется только name(Имя автора)
	//
	// in:body
	Body model.Book
}


// swagger:route DELETE /authors/{author_id}/books/{book_id} books bookDeleteRequest
// Удаление книги автора
// responses:
//   200: bookEmptyResponse

// swagger:parameters bookDeleteRequest
type bookDeleteRequest struct {
	// ID автора
	//
	// in:path
	ID_author string `json:"author_id"`
	// ID книги
	//
	// in:path
	ID_book string `json:"book_id"`
}