package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type FieldRepository interface {
	Save(ctx context.Context, tx *gorm.DB, field *models.Field) error
	Update(ctx context.Context, tx *gorm.DB, field *models.Field) error
	Delete(ctx context.Context, tx *gorm.DB, fieldId uint) error
	FindById(ctx context.Context, tx *gorm.DB, fieldId uint) (*models.Field, error)
	FindAll(ctx context.Context, tx *gorm.DB) (*[]models.Field, error)
}
