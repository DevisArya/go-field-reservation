package models

import (
	"time"
)

type Transaction struct {
	Id                uint `gorm:"primary_key"`
	UserId            uint
	PaymentUrl        string
	OrderId           string
	TotalPrice        int64
	TransactionTime   time.Time
	PaymentType       string
	PaymentStatus     string
	SettlementTime    time.Time
	FraudStatus       string
	TransactionDetail []TransactionDetail `gorm:"foreignKey:TransactionId"`
}
