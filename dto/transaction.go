package dto

import "time"

type TransactionDetailReq struct {
	ScheduleId uint  `json:"ScheduleId" form:"ScheduleId" validate:"required"`
	Price      int64 `json:"Price" form:"Price" validate:"required"`
}

type TransactionReq struct {
	TransactionDetail []TransactionDetailReq `json:"TransactionDetail" form:"TransactionDetail" validate:"required"`
}

type TransactionCreateResponse struct {
	Id            uint
	PaymentStatus string
	PaymentUrl    string
}

type UpdateTransactionResponse struct {
	Id            uint
	PaymentStatus string
	UpdatedAt     string
}

type TransactionDetailResponse struct {
	Id            uint
	TransactionId uint
	ScheduleId    uint
	Price         int64
}

type TransactionResponse struct {
	Id                uint
	UserId            uint
	PaymentUrl        string
	PaymentStatus     string
	TotalPrice        int64
	CreatedAt         time.Time
	CanceledReason    string
	TransactionDetail []TransactionDetailResponse
}
