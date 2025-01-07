package service

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/dto"
)

type TransactionService interface {
	Save(ctx context.Context, transactionData *dto.TransactionReq, userId uint) (*dto.TransactionCreateResponse, error)
	Update(ctx context.Context, transactionData *dto.MidtransRequest) error
}
