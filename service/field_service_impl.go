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

type FieldServiceImpl struct {
	FieldRepository repository.FieldRepository
	DB              *gorm.DB
	validate        *validator.Validate
}

func NewFieldService(FieldRepository repository.FieldRepository, DB *gorm.DB, validate *validator.Validate) FieldService {
	return &FieldServiceImpl{
		FieldRepository,
		DB,
		validate,
	}
}

// Save implements FieldService
func (service *FieldServiceImpl) Save(ctx context.Context, request *dto.FieldReqRes) (*dto.FieldReqRes, error) {

	if err := service.validate.Struct(request); err != nil {
		return nil, err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	fieldData := models.Field{
		Name:  request.Name,
		Type:  request.Type,
		Price: request.Price,
	}

	if err := service.FieldRepository.Save(ctx, tx, &fieldData); err != nil {
		return nil, err
	}

	return request, nil
}

// Update implements FieldService
func (service *FieldServiceImpl) Update(ctx context.Context, request *models.Field) (*models.Field, error) {

	if err := service.validate.Struct(request); err != nil {
		return nil, err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.FieldRepository.FindById(ctx, tx, request.Id); err != nil {
		return nil, err
	}

	if err := service.FieldRepository.Update(ctx, tx, request); err != nil {
		return nil, err
	}

	return request, nil
}

// Delete implements FieldService
func (service *FieldServiceImpl) Delete(ctx context.Context, fieldId uint) error {

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.FieldRepository.FindById(ctx, tx, fieldId); err != nil {
		return err
	}

	if err := service.FieldRepository.Delete(ctx, tx, fieldId); err != nil {
		return err
	}

	return nil
}

// FindById implements FieldService
func (service *FieldServiceImpl) FindById(ctx context.Context, fieldId uint) (*models.Field, error) {

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	field, err := service.FieldRepository.FindById(ctx, tx, fieldId)
	if err != nil {
		return nil, err
	}

	return field, nil
}

// FindAll implements FieldService
func (service *FieldServiceImpl) FindAll(ctx context.Context) (*[]models.Field, error) {

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	fields, err := service.FieldRepository.FindAll(ctx, tx)

	if err != nil {
		return nil, err
	}

	return fields, nil

}
