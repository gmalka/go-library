package apiPresentations

import "go-library/model"

// swagger:parameters userEmptyRequest
type userEmptyRequest struct {}

// swagger:response userEmptyResponse
type userEmptyResponse struct {}


// swagger:route GET /users users userEmptyRequest
// Получение всех пользователей
// responses:
//   200: userGetAllResponse

// swagger:response userGetAllResponse
type userGetAllResponse struct {
	// Массив юзеров
	//
	// in:body
	Body []model.User
}


// swagger:route POST /users users userAddRequest
// Создание нового пользователя
// responses:
//   200: userEmptyResponse

// swagger:parameters userAddRequest
type userAddRequest struct {
	// Требуется только name(Имя пользователя)
	//
	// in:body
	Body model.User
}


// swagger:route DELETE /users/{user_id} users userDeleteRequest
// Удаление пользователя
// responses:
//   200: userEmptyResponse

// swagger:parameters userDeleteRequest
type userDeleteRequest struct {
	// ID удаляемого пользователя
	//
	// in:path
	ID string `json:"user_id"`
}