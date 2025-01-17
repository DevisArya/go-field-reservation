package app

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	"github.com/DevisArya/reservasi_lapangan/repository"
	"github.com/DevisArya/reservasi_lapangan/service"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NewAppContainer(db *gorm.DB, validate *validator.Validate) *handler.AppHandler {

	// inisialisasi repositories
	userRepository := repository.NewUserRepository()
	operatorRepository := repository.NewOperatorRepository()
	fieldRepository := repository.NewFieldRepository()
	scheduleRepository := repository.NewScheduleRepository()
	transactionRepository := repository.NewTransactionRepository()
	authRepository := repository.NewAuthRepository()

	//inisialisasi services
	userService := service.NewUserService(userRepository, db, validate)
	operatorService := service.NewOperatorService(operatorRepository, db, validate)
	fieldService := service.NewFieldService(fieldRepository, db, validate)
	scheduleService := service.NewScheduleService(scheduleRepository, db, validate)
	transactionService := service.NewTransactionService(transactionRepository, db, validate)
	authService := service.NewAuthService(authRepository, db, validate)

	//inisialisasi handlers
	userHandler := handler.NewUserHandler(userService)
	operatorHandler := handler.NewOperatorHandler(operatorService)
	fieldHandler := handler.NewFieldHandler(fieldService)
	scheduleHandler := handler.NewScheduleHandler(scheduleService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	authHandler := handler.NewAuthHandler(authService)

	return &handler.AppHandler{
		UserHandler:        userHandler,
		OperatorHandler:    operatorHandler,
		FieldHandler:       fieldHandler,
		ScheduleHandler:    scheduleHandler,
		TransactionHandler: transactionHandler,
		AuthHandler:        authHandler,
	}
}
