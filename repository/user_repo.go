package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, tx *gorm.DB, user *models.User) error
	Update(ctx context.Context, tx *gorm.DB, user *models.User) error
	Delete(ctx context.Context, tx *gorm.DB, userId uint) error
	FindById(ctx context.Context, tx *gorm.DB, userId uint) (*models.User, error)
	FindAll(ctx context.Context, tx *gorm.DB) (*[]models.User, error)
}
