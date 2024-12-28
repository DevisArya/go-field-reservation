package service

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/dto"
	"github.com/DevisArya/reservasi_lapangan/models"
)

type UserService interface {
	Create(ctx context.Context, request *dto.UserCreateRequest) error
	Update(ctx context.Context, request *models.User) error
	Delete(ctx context.Context, id uint) error
	FindById(ctx context.Context, id uint) (*models.User, error)
	FindAll(ctx context.Context) (*[]models.User, error)
}
