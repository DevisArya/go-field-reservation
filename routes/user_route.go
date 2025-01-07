package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	mdlw "github.com/DevisArya/reservasi_lapangan/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo, userHandler handler.UserHandler) {

	// route for user
	e.POST("/user", userHandler.Create)
	e.PATCH("/user/:id", userHandler.Update, mdlw.Auth([]string{"user"}, true))
	e.DELETE("/user/:id", userHandler.Delete, mdlw.Auth([]string{"user,"}, true))
	e.GET("/user/:id", userHandler.FindById, mdlw.Auth([]string{"user"}, true))
	e.GET("/users", userHandler.FindAll, mdlw.Auth([]string{"super user", "operator"}, false))
}
