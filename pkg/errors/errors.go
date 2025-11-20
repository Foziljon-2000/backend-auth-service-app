package errors

import "errors"

var (
	ErrInternalServer   = errors.New("Внутреняя ошибка сервера")
	ErrBadRequest       = errors.New("Неправильный запрос")
	ErrMethodNotAllowed = errors.New("Неправильный метод запроса")
	ErrSuccess          = errors.New("Успешно")
)

var StatusCodes map[string]int = map[string]int{
	// 200 code
	ErrSuccess.Error(): 200,

	// 400 code
	ErrBadRequest.Error(): 400,

	// 500 code
	ErrInternalServer.Error():   500,
	ErrMethodNotAllowed.Error(): 502,
}
