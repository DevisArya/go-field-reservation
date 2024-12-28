package service

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/dto"
	"github.com/DevisArya/reservasi_lapangan/helper"
	"github.com/DevisArya/reservasi_lapangan/models"
	"github.com/DevisArya/reservasi_lapangan/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OperatorServiceImpl struct {
	OperatorRepository repository.OperatorRepository
	DB                 *gorm.DB
	validate           *validator.Validate
}

func NewOperatorService(operatorRepository repository.OperatorRepository, DB *gorm.DB, validate *validator.Validate) OperatorService {
	return &OperatorServiceImpl{
		OperatorRepository: operatorRepository,
		DB:                 DB,
		validate:           validate,
	}
}

// Create implements OperatorService
func (service *OperatorServiceImpl) Create(ctx context.Context, request *dto.OperatorCreateRequest) error {
	if err := service.validate.Struct(request); err != nil {
		return err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	operatorData := models.Operator{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
	}

	if err := service.OperatorRepository.Save(ctx, tx, &operatorData); err != nil {
		return err
	}
	return nil
}

// Update implements OperatorService
func (service *OperatorServiceImpl) Update(ctx context.Context, request *models.Operator) error {
	if err := service.validate.Struct(request); err != nil {
		return err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.OperatorRepository.FindById(ctx, tx, request.Id); err != nil {
		return err
	}

	if err := service.OperatorRepository.Update(ctx, tx, request); err != nil {
		return err
	}

	return nil
}

// Delete implements OperatorService
func (service *OperatorServiceImpl) Delete(ctx context.Context, id uint) error {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.OperatorRepository.FindById(ctx, tx, id); err != nil {
		return err
	}

	if err := service.OperatorRepository.Delete(ctx, tx, id); err != nil {
		return err
	}

	return nil
}

// FindById implements OperatorService
func (service *OperatorServiceImpl) FindById(ctx context.Context, id uint) (*models.Operator, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	operator, err := service.OperatorRepository.FindById(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return operator, nil
}

// FindAll implements OperatorService
func (service *OperatorServiceImpl) FindAll(ctx context.Context) (*[]models.Operator, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	operators, err := service.OperatorRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return operators, nil
}
