package service

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"github.com/DevisArya/reservasi_lapangan/models/web"
)

type UserService interface {
	Create(ctx context.Context, request *web.UserCreateRequest) error
	Update(ctx context.Context, request *models.User) error
	Delete(ctx context.Context, id uint) error
	FindById(ctx context.Context, id uint) (*models.User, error)
	FindAll(ctx context.Context) (*[]models.User, error)
}
