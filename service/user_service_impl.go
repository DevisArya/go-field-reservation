package service

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/helper"
	"github.com/DevisArya/reservasi_lapangan/models"
	"github.com/DevisArya/reservasi_lapangan/models/web"
	"github.com/DevisArya/reservasi_lapangan/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		validate:       validate,
	}
}

// Create implements UserService
func (service *UserServiceImpl) Create(ctx context.Context, request *web.UserCreateRequest) error {
	if err := service.validate.Struct(request); err != nil {
		return err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	userData := models.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
	}

	if err := service.UserRepository.Save(ctx, tx, &userData); err != nil {
		return err
	}
	return nil
}

// Update implements UserService
func (service *UserServiceImpl) Update(ctx context.Context, request *models.User) error {
	if err := service.validate.Struct(request); err != nil {
		return err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.UserRepository.FindById(ctx, tx, request.Id); err != nil {
		return err
	}

	if err := service.UserRepository.Update(ctx, tx, request); err != nil {
		return err
	}

	return nil
}

// Delete implements UserService
func (service *UserServiceImpl) Delete(ctx context.Context, id uint) error {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.UserRepository.FindById(ctx, tx, id); err != nil {
		return err
	}

	if err := service.UserRepository.Delete(ctx, tx, id); err != nil {
		return err
	}

	return nil
}

// FindById implements UserService
func (service *UserServiceImpl) FindById(ctx context.Context, id uint) (*models.User, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindAll implements UserService
func (service *UserServiceImpl) FindAll(ctx context.Context) (*[]models.User, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	users, err := service.UserRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
