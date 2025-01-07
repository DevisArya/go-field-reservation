package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, AppHandler *handler.AppHandler) {

	// Regist All Routes
	RegisterUserRoutes(e, AppHandler.UserHandler)
	RegisterFieldRoutes(e, AppHandler.FieldHandler)
	RegisterOperatorRoutes(e, AppHandler.OperatorHandler)
	RegisterScheduleRoutes(e, AppHandler.ScheduleHandler)
	RegisterTransactionRoutes(e, AppHandler.TransactionHandler)
	RegisterAuthRoutes(e, AppHandler.AuthHandler)
}
