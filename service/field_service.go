package service

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/dto"
	"github.com/DevisArya/reservasi_lapangan/models"
)

type FieldService interface {
	Save(ctx context.Context, request *dto.FieldReqRes) (*dto.FieldReqRes, error)
	Update(ctx context.Context, request *models.Field) (*models.Field, error)
	Delete(ctx context.Context, fieldId uint) error
	FindById(ctx context.Context, fieldId uint) (*models.Field, error)
	FindAll(ctx context.Context) (*[]models.Field, error)
}
