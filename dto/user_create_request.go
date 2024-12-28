package dto

type UserCreateRequest struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Name     string `json:"Name" form:"Name" validate:"required"`
	Password string `json:"Password" form:"Password" validate:"required"`
}
