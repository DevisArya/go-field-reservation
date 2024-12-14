package models

import "time"

type Reservations struct {
	Id         uint      `gorm:"primary_key"`
	UserId     uint      `json:"UserId " form:"UserId" validate:"required"`
	Status     string    `json:"Status" form:"Status" validate:"required"`
	Start_time time.Time `validate:"required"`
	Lenght     uint      `json:"Lenght" validate:"required"`
	CreatedAt  time.Time `validate:"required"`
}
