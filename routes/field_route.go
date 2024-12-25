package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	"github.com/labstack/echo/v4"
)

func RegisterFieldRoutes(e *echo.Echo, fieldHandler handler.FieldHandler) {

	// route for user
	e.POST("/field", fieldHandler.Create)
	e.PATCH("/field/:id", fieldHandler.Update)
	e.DELETE("/field/:id", fieldHandler.Delete)
	e.GET("/field/:id", fieldHandler.FindById)
	e.GET("/fields", fieldHandler.FindAll)
}
