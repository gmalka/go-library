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