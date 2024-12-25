package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type FieldRepositoryImpl struct{}

func NewFieldRepository() FieldRepository {
	return &FieldRepositoryImpl{}
}

// Save implements FieldRepository
func (repository *FieldRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, field *models.Field) error {

	if err := tx.WithContext(ctx).Create(field).Error; err != nil {
		return err
	}
	return nil
}

// Update implements FieldRepository
func (repository *FieldRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, field *models.Field) error {

	if err := tx.WithContext(ctx).Updates(&field).Error; err != nil {
		return err
	}

	return nil
}

// Delete implements FieldRepository
func (repository *FieldRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, fieldId uint) error {

	if err := tx.WithContext(ctx).Delete(&models.Field{}, fieldId).Error; err != nil {
		return err
	}

	return nil
}

// FindById implements FieldRepository
func (repository *FieldRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, fieldId uint) (*models.Field, error) {
	var field models.Field

	if err := tx.WithContext(ctx).First(&field, fieldId).Error; err != nil {
		return nil, err
	}
	return &field, nil
}

// FindAll implements FieldRepository
func (repository *FieldRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) (*[]models.Field, error) {
	var fields []models.Field

	if err := tx.WithContext(ctx).Find(&fields).Error; err != nil {
		return nil, err
	}

	return &fields, nil
}
