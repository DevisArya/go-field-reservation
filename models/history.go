package models

import (
	"time"
)

type History struct {
	Id         uint
	OperatorId uint
	FieldId    uint
	Time       time.Time
	Date       time.Time
}
