package models

import (
	"time"
)

type Transaction struct {
	Id                uint `gorm:"primary_key"`
	TransactionId     string
	UserId            uint
	PaymentUrl        string
	OrderId           string
	TotalPrice        int64
	TransactionTime   time.Time `gorm:"column:transaction_time;type:datetime;null"`
	PaymentType       string
	PaymentStatus     string
	SettlementTime    time.Time `gorm:"column:settlement_time;type:datetime;null"`
	FraudStatus       string
	TransactionDetail []TransactionDetail `gorm:"foreignKey:TransactionId"`
}
