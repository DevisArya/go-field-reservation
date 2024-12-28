package service

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
)

type ScheduleService interface {
	Create(ctx context.Context, request *models.Schedule) error
	Update(ctx context.Context, request *models.Schedule) error
	Delete(ctx context.Context, id uint) error
	FindById(ctx context.Context, id uint) (*models.Schedule, error)
	FindAll(ctx context.Context) (*[]models.Schedule, error)
}
