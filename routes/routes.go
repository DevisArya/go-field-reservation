package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, userHandler handler.UserHandler) *echo.Echo {

	g_user := e.Group("/user")
	g_user.POST("", userHandler.Create)
	g_user.PUT("/:id", userHandler.Update)
	g_user.DELETE("/:id", userHandler.Delete)
	g_user.GET("/:id", userHandler.FindById)
	g_user.GET("s", userHandler.FindAll)

	return e
}
