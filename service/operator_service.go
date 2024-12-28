package service

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/dto"
	"github.com/DevisArya/reservasi_lapangan/models"
)

type OperatorService interface {
	Create(ctx context.Context, request *dto.OperatorCreateRequest) error
	Update(ctx context.Context, request *models.Operator) error
	Delete(ctx context.Context, id uint) error
	FindById(ctx context.Context, id uint) (*models.Operator, error)
	FindAll(ctx context.Context) (*[]models.Operator, error)
}
