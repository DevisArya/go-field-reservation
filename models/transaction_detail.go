package models

type TransactionDetail struct {
	Id            uint `gorm:"primary_key"`
	TransactionId uint
	ScheduleId    uint
	Name          string
	Price         int64
}
