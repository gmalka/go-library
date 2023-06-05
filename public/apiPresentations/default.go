package apiPresentations

import "go-library/model"


// swagger:route GET / default defaultGetAllTakenBooksRequest
// Получение взятых книг и взявших их пользователей
// responses:
//   200: defaultGetAllResponse

// swagger:parameters defaultEmptyRequest
type defaultGetAllTakenBooksRequest struct {}

// swagger:response defaultGetAllResponse
type defaultGetAllResponse struct {
	// Массив книг и взявших их книг
	//
	// in:body
	Body []model.TakenBook
}


// swagger:route GET /books default defaultGetAllBooksRequest
// Получение всех книг
// responses:
//   200: defaultGetAllBooksResponse

// swagger:parameters defaultEmptyRequest
type defaultGetAllBooksRequest struct {}

// swagger:response defaultGetAllBooksResponse
type defaultGetAllBooksResponse struct {
	// Массив книг
	//
	// in:body
	Body []model.BookWithAuthor
}