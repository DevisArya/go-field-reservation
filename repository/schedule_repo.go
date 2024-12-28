package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Save(ctx context.Context, tx *gorm.DB, schedule *models.Schedule) error
	Update(ctx context.Context, tx *gorm.DB, schedule *models.Schedule) error
	Delete(ctx context.Context, tx *gorm.DB, scheduleId uint) error
	FindById(ctx context.Context, tx *gorm.DB, scheduleId uint) (*models.Schedule, error)
	FindAll(ctx context.Context, tx *gorm.DB) (*[]models.Schedule, error)
}
