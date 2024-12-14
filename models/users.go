package models

type User struct {
	Id       uint   `json:"Id" form:"Id" gorm:"primary_key;not null"`
	Email    string `json:"Email" form:"Email" validate:"required,email" gorm:"not null;type:VARCHAR(100);uniqueIndex"`
	Name     string `json:"Name" form:"Name" validate:"required" gorm:"type:VARCHAR(50);not null"`
	Password string `json:"Password" form:"Password" validate:"required" gorm:"type:VARCHAR(255);not null"`
}
