package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

// Save implements TransactionRepository
func (*TransactionRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, transactionData *models.Transaction) (*models.Transaction, error) {

	if err := tx.WithContext(ctx).Create(transactionData).Error; err != nil {
		return nil, err
	}

	return transactionData, nil
}

// Update implements TransactionRepository
func (*TransactionRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, updateData *models.Transaction) error {

	if err := tx.WithContext(ctx).Where("transaction_id = ?", updateData.TransactionId).Updates(&updateData).Error; err != nil {
		return err
	}

	return nil

}
