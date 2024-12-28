package repository

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

type ScheduleRepositoryImpl struct {
}

func NewScheduleRepository() ScheduleRepository {
	return &ScheduleRepositoryImpl{}
}

// Save implements ScheduleRepository
func (*ScheduleRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, schedule *models.Schedule) error {

	if err := tx.WithContext(ctx).Create(schedule).Error; err != nil {
		return err
	}

	return nil
}

// Update implements ScheduleRepository
func (*ScheduleRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, schedule *models.Schedule) error {

	if err := tx.WithContext(ctx).Updates(&schedule).Error; err != nil {
		return err
	}

	return nil
}

// Delete implements ScheduleRepository
func (*ScheduleRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, scheduleId uint) error {

	if err := tx.WithContext(ctx).Delete(&models.Schedule{}, scheduleId).Error; err != nil {
		return err
	}

	return nil
}

// FindById implements ScheduleRepository
func (*ScheduleRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, scheduleId uint) (*models.Schedule, error) {
	var schedule models.Schedule

	if err := tx.WithContext(ctx).First(&schedule, scheduleId).Error; err != nil {
		return nil, err
	}

	return &schedule, nil
}

// FindAll implements ScheduleRepository
func (*ScheduleRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) (*[]models.Schedule, error) {
	var schedules []models.Schedule

	if err := tx.WithContext(ctx).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return &schedules, nil
}
