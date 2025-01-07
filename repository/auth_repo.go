package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByEmail(ctx context.Context, tx *gorm.DB, email string) (*models.User, error)
}
