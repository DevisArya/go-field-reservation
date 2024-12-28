package models

import (
	"time"
)

type Transaction struct {
	Id                uint `gorm:"primary_key"`
	UserId            uint `json:"UserId " form:"UserId" validate:"required"`
	PaymentUrl        string
	PaymentStatus     string
	TotalPrice        uint64
	CreatedAt         time.Time
	CanceledReason    string
	TransactionDetail string
}
