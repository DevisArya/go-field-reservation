package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

// FindByEmail implements AuthRepository
func (*AuthRepositoryImpl) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (*models.User, error) {

	var user models.User

	if err := tx.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
