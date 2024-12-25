package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo, userHandler handler.UserHandler) {

	// route for user
	e.POST("/user", userHandler.Create)
	e.PATCH("/user/:id", userHandler.Update)
	e.DELETE("/user/:id", userHandler.Delete)
	e.GET("/user/:id", userHandler.FindById)
	e.GET("/users", userHandler.FindAll)
}
