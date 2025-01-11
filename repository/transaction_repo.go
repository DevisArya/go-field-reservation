package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(ctx context.Context, tx *gorm.DB, transactionData *models.Transaction) (*models.Transaction, error)
	Update(ctx context.Context, tx *gorm.DB, transactionData *models.Transaction) error
	GetScheduleByIds(ctx context.Context, tx *gorm.DB, id []uint) (int, error)
	LockScheduleAndUpdate(ctx context.Context, tx *gorm.DB, scheduleIds []uint) error
	GetTransactionById(ctx context.Context, tx *gorm.DB, id string) (*models.Transaction, error)
	UpdateSchedulesStatus(ctx context.Context, tx *gorm.DB, scheduleIds []uint, status string) error
}
