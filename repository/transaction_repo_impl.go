package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

// Save implements TransactionRepository
func (*TransactionRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, transactionData *models.Transaction) (*models.Transaction, error) {

	if err := tx.WithContext(ctx).
		Create(transactionData).
		Error; err != nil {
		return nil, err
	}

	return transactionData, nil
}

// Update implements TransactionRepository
func (*TransactionRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, updateData *models.Transaction) error {

	if err := tx.WithContext(ctx).
		Where("transaction_id = ?", updateData.TransactionId).
		Updates(&updateData).
		Error; err != nil {
		return err
	}

	return nil

}

func (*TransactionRepositoryImpl) GetScheduleByIds(ctx context.Context, tx *gorm.DB, ids []uint) (int, error) {

	var schedules []models.Schedule

	if err := tx.WithContext(ctx).
		Where("id IN ?", ids).
		Where("status = ?", "available").
		Find(&schedules).
		Error; err != nil {
		return 0, err
	}

	return len(schedules), nil
}

// Update Available Schedule implements Transaction Repository

func (*TransactionRepositoryImpl) LockScheduleAndUpdate(ctx context.Context, tx *gorm.DB, scheduleIds []uint) error {

	if err := tx.WithContext(ctx).
		Model(&models.Schedule{}).
		Where("id IN ?", scheduleIds).
		Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
		Updates(map[string]interface{}{"status": "reserved"}).
		Error; err != nil {
		return err
	}

	return nil
}

func (*TransactionRepositoryImpl) GetTransactionById(ctx context.Context, tx *gorm.DB, id string) (*models.Transaction, error) {
	var transaction models.Transaction

	if err := tx.WithContext(ctx).
		Where("transaction_id = ?", id).
		First(&transaction).
		Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

// Update Status Schedule implements Transaction Repository
func (*TransactionRepositoryImpl) UpdateSchedulesStatus(ctx context.Context, tx *gorm.DB, scheduleIds []uint, status string) error {

	if err := tx.WithContext(ctx).
		Model(&models.Schedule{}).
		Where("id IN ?", scheduleIds).
		Updates(map[string]interface{}{"status": status}).
		Error; err != nil {
		return err
	}
	return nil
}
