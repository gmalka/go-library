package apiPresentations

import "go-library/model"

// swagger:parameters libraryEmptyRequest
type libraryEmptyRequest struct {}

// swagger:response libraryEmptyResponse
type libraryEmptyResponse struct {}



// swagger:route GET /users/{user_id}/library library getTakenBookRequest
// Получение всех взятых юзером книг
// responses:
//   200: getTakenBookResponse

// swagger:parameters getTakenBookRequest
type getTakenBookRequest struct {
	// ID юзера
	//
	// in:path
	ID_author string `json:"user_id"`
}

// swagger:response getTakenBookResponse
type getTakenBookResponse struct {
	// Массив книг и их авторов
	//
	// in:body
	Body []model.BookWithAuthor
}



// swagger:route POST /users/{user_id}/library/{book_id} library takeBookRequest
// Взятие новой книги пользователем
// responses:
//   200: libraryEmptyResponse

// swagger:parameters takeBookRequest
type takeBookRequest struct {
	// ID юзера
	//
	// in:path
	ID_author string `json:"user_id"`
	// ID книги
	//
	// in:path
	ID_book string `json:"book_id"`
}


// swagger:route DELETE /users/{user_id}/library/{book_id} library putBackBookRequest
// Возвращение взятой книги пользователем
// responses:
//   200: libraryEmptyResponse

// swagger:parameters putBackBookRequest
type putBackBookRequest struct {
	// ID юзера
	//
	// in:path
	ID_author string `json:"user_id"`
	// ID книги
	//
	// in:path
	ID_book string `json:"book_id"`
}