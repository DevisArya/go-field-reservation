package models

type Field struct {
	Id    uint   `gorm:"primary_key"`
	Name  string `json:"Name" form:"Name" validate:"required"`
	Type  string `json:"Type" form:"Type" validate:"requierd"`
	Price uint   `json:"Price" form:"Price" validate:"required"`
}
