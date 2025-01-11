package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// Save implements UserRepository
func (*UserRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, user *models.User) error {

	if err := tx.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	return nil
}

// Update implements UserRepository
func (*UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user *models.User) error {

	if err := tx.WithContext(ctx).Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

// Delete implements UserRepository
func (*UserRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, userId uint) error {

	if err := tx.WithContext(ctx).Delete(&models.User{}, userId).Error; err != nil {
		return err
	}

	return nil
}

// FindById implements UserRepository
func (*UserRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, userId uint) (*models.User, error) {
	var user models.User

	if err := tx.WithContext(ctx).First(&user, userId).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAll implements UserRepository
func (*UserRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) (*[]models.User, error) {
	var users []models.User

	if err := tx.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
