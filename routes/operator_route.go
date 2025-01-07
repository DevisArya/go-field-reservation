package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	mdlw "github.com/DevisArya/reservasi_lapangan/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterOperatorRoutes(e *echo.Echo, operatorHandler handler.OperatorHandler) {

	// route for operator
	operator := e.Group("/operator", mdlw.Auth([]string{"super user"}, false))
	operator.POST("", operatorHandler.Create)
	operator.PATCH("/:id", operatorHandler.Update)
	operator.DELETE("/:id", operatorHandler.Delete)
	operator.GET("/:id", operatorHandler.FindById)
	operator.GET("s", operatorHandler.FindAll)
}
