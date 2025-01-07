package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	Id       uint
	Username string
	Role     string
}

type LoginResponse struct {
	User  UserLoginResponse
	Token string
}
