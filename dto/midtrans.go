package dto

import "time"

type MidtransRequest struct {
	StatusCode        string    `json:"status_code"`
	TransactionTime   time.Time `json:"transaction_time"`
	TransactionId     string    `json:"transaction_id"`
	TransactionStatus string    `json:"transaction_status"`
	StatusMessage     string    `json:"status_messatime"`
	SignatureKey      string    `json:"signature_key"`
	SettlementTime    time.Time `json:"settlement_time"`
	PaymentType       string    `json:"payment_type"`
	OrderId           string    `json:"order_id"`
	GrossAmount       string    `json:"gross_amount"`
	FraudStatus       string    `json:"fraud_status"`
}
