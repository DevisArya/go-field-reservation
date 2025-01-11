package models

import (
	"time"
)

type Schedule struct {
	Id         uint
	OperatorId uint
	FieldId    uint
	Time       time.Time
	Date       time.Time
	Status     string
}
