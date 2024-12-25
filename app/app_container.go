package app

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	"github.com/DevisArya/reservasi_lapangan/repository"
	"github.com/DevisArya/reservasi_lapangan/service"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// type AppContainer struct {
// 	UserHandler  handler.UserHandler
// 	FieldHandler handler.FieldHandler
// }

func NewAppContainer(db *gorm.DB, validate *validator.Validate) *handler.AppHandler {

	// inisialisasi repositories
	userRepository := repository.NewUserRepository()
	fieldRepository := repository.NewFieldRepository()

	//inisialisasi services
	userService := service.NewUserService(userRepository, db, validate)
	fieldService := service.NewFieldService(fieldRepository, db, validate)

	//inisialisasi handlers
	userHandler := handler.NewUserHandler(userService)
	fieldHandler := handler.NewFieldHandler(fieldService)

	return &handler.AppHandler{
		UserHandler:  userHandler,
		FieldHandler: fieldHandler,
	}
}
