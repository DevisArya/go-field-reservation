package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	"github.com/labstack/echo/v4"
)

func RegisterOperatorRoutes(e *echo.Echo, operatorHandler handler.OperatorHandler) {

	// route for operator
	e.POST("/operator", operatorHandler.Create)
	e.PATCH("/operator/:id", operatorHandler.Update)
	e.DELETE("/operator/:id", operatorHandler.Delete)
	e.GET("/operator/:id", operatorHandler.FindById)
	e.GET("/operators", operatorHandler.FindAll)
}
