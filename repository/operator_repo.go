package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type OperatorRepository interface {
	Save(ctx context.Context, tx *gorm.DB, operator *models.Operator) error
	Update(ctx context.Context, tx *gorm.DB, operator *models.Operator) error
	Delete(ctx context.Context, tx *gorm.DB, operatorId uint) error
	FindById(ctx context.Context, tx *gorm.DB, operatorId uint) (*models.Operator, error)
	FindAll(ctx context.Context, tx *gorm.DB) (*[]models.Operator, error)
}
