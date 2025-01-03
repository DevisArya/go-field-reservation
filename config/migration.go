package config

import (
	"github.com/DevisArya/reservasi_lapangan/helper"
	"github.com/DevisArya/reservasi_lapangan/models"
	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {

	err := db.AutoMigrate(
		&models.User{},
		&models.Field{},
		&models.Operator{},
		&models.Schedule{},
		&models.Transaction{},
		&models.TransactionDetail{},
	)

	helper.PanicIfError(err)
}
