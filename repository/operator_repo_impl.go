package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type OperatorRepositoryImpl struct {
}

func NewOperatorRepository() OperatorRepository {
	return &OperatorRepositoryImpl{}
}

// Save implements OperatorRepository
func (*OperatorRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, operator *models.Operator) error {

	if err := tx.WithContext(ctx).Create(operator).Error; err != nil {
		return err
	}

	return nil
}

// Update implements OperatorRepository
func (*OperatorRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, operator *models.Operator) error {

	if err := tx.WithContext(ctx).Where("id = ?", operator.Id).Updates(&operator).Error; err != nil {
		return err
	}

	return nil
}

// Delete implements OperatorRepository
func (*OperatorRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, operatorId uint) error {

	if err := tx.WithContext(ctx).Delete(&models.Operator{}, operatorId).Error; err != nil {
		return err
	}

	return nil
}

// FindById implements OperatorRepository
func (*OperatorRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, operatorId uint) (*models.Operator, error) {
	var operator models.Operator

	if err := tx.WithContext(ctx).First(&operator, operatorId).Error; err != nil {
		return nil, err
	}

	return &operator, nil
}

// FindAll implements OperatorRepository
func (*OperatorRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) (*[]models.Operator, error) {
	var operators []models.Operator

	if err := tx.WithContext(ctx).Find(&operators).Error; err != nil {
		return nil, err
	}
	return &operators, nil
}
