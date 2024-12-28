package repository

import (
	"context"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, tx gorm.DB)
}
