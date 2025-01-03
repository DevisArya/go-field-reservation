package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(ctx context.Context, tx *gorm.DB, transactionData *models.Transaction) (*models.Transaction, error)
	Update(ctx context.Context, tx *gorm.DB, transactionData *models.Transaction) error
}
