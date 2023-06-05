package apiPresentations

import "go-library/model"

// swagger:parameters authorEmptyRequest
type vacancyAddRequest struct {}

// swagger:response authorEmptyResponse
type authorEmptyRequest struct {}

// swagger:route GET /authors authors authorEmptyRequest
// Получение всех авторов
// responses:
//   200: authorGetAllResponse

// swagger:response authorGetAllResponse
type authorGetAllResponse struct {
	// Массив авторов
	//
	// in:body
	Body []model.Author
}


// swagger:route POST /authors authors authorAddRequest
// Создание нового автора
// responses:
//   200: authorEmptyResponse

// swagger:parameters authorAddRequest
type authorAddRequest struct {
	// Достаточно только поля name(Имя автора)
	//
	// in:body
	Body model.Author
}


// swagger:route POST /authors/{author_id} authors authorDeleteRequest
// Удаление автора
// responses:
//   200: authorEmptyResponse

// swagger:parameters authorDeleteRequest
type authorDeleteRequest struct {
	// ID удаляемого автора
	//
	// in:path
	ID string `json:"author_id"`
}