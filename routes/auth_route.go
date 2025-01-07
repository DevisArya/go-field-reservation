package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo, authHandler handler.AuthHandler) {

	//route for auth
	e.POST("/login", authHandler.Login)
}
