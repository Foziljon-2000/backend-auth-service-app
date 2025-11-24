package handler

type CreateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
