package errors

import "errors"

var (
	ErrInternalServer         = errors.New("Внутреняя ошибка сервера")
	ErrBadRequest             = errors.New("Неправильный запрос")
	ErrMethodNotAllowed       = errors.New("Неправильный метод запроса")
	ErrSuccess                = errors.New("Успешно")
	ErrUserDoesNotExist       = errors.New("Такого пользователя не существует")
	ErrUserLoginAlreadyExists = errors.New("Такой логин уже существует")
	ErrAccountDoesNotExist    = errors.New("Аккаунта не существует")
)

var StatusCodes map[string]int = map[string]int{
	// 200 code
	ErrSuccess.Error(): 200,

	// 400 code
	ErrBadRequest.Error():             400,
	ErrUserDoesNotExist.Error():       404,
	ErrAccountDoesNotExist.Error():    404,
	ErrUserLoginAlreadyExists.Error(): 409,

	// 500 code
	ErrInternalServer.Error():   500,
	ErrMethodNotAllowed.Error(): 502,
}
