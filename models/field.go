package models

type Field struct {
	Id    uint   `gorm:"primary_key"`
	Name  string `json:"Name" form:"Name" validate:"omitempty"`
	Type  string `json:"Type" form:"Type" validate:"omitempty"`
	Price uint   `json:"Price" form:"Price" validate:"omitempty"`
}
