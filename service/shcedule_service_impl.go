package service

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/helper"
	"github.com/DevisArya/reservasi_lapangan/models"
	"github.com/DevisArya/reservasi_lapangan/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ScheduleServiceImpl struct {
	ScheduleRepository repository.ScheduleRepository
	DB                 *gorm.DB
	validate           *validator.Validate
}

func NewScheduleService(scheduleRepository repository.ScheduleRepository, DB *gorm.DB, validate *validator.Validate) ScheduleService {
	return &ScheduleServiceImpl{
		ScheduleRepository: scheduleRepository,
		DB:                 DB,
		validate:           validate,
	}
}

// Create implements ScheduleService
func (service *ScheduleServiceImpl) Create(ctx context.Context, request *models.Schedule) error {
	if err := service.validate.Struct(request); err != nil {
		return err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if err := service.ScheduleRepository.Save(ctx, tx, request); err != nil {
		return err
	}
	return nil
}

// Update implements ScheduleService
func (service *ScheduleServiceImpl) Update(ctx context.Context, request *models.Schedule) error {
	if err := service.validate.Struct(request); err != nil {
		return err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.ScheduleRepository.FindById(ctx, tx, request.Id); err != nil {
		return err
	}

	if err := service.ScheduleRepository.Update(ctx, tx, request); err != nil {
		return err
	}

	return nil
}

// Delete implements ScheduleService
func (service *ScheduleServiceImpl) Delete(ctx context.Context, id uint) error {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	if _, err := service.ScheduleRepository.FindById(ctx, tx, id); err != nil {
		return err
	}

	if err := service.ScheduleRepository.Delete(ctx, tx, id); err != nil {
		return err
	}

	return nil
}

// FindById implements ScheduleService
func (service *ScheduleServiceImpl) FindById(ctx context.Context, id uint) (*models.Schedule, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	schedule, err := service.ScheduleRepository.FindById(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

// FindAll implements ScheduleService
func (service *ScheduleServiceImpl) FindAll(ctx context.Context) (*[]models.Schedule, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	schedules, err := service.ScheduleRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return schedules, nil
}
