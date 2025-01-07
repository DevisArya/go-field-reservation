package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	mdlw "github.com/DevisArya/reservasi_lapangan/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterFieldRoutes(e *echo.Echo, fieldHandler handler.FieldHandler) {

	// route for user
	e.POST("/field", fieldHandler.Create, mdlw.Auth([]string{"operator, super user"}, false))
	e.PATCH("/field/:id", fieldHandler.Update, mdlw.Auth([]string{"operator, super user"}, false))
	e.DELETE("/field/:id", fieldHandler.Delete, mdlw.Auth([]string{"operator, super user"}, false))
	e.GET("/field/:id", fieldHandler.FindById)
	e.GET("/fields", fieldHandler.FindAll)
}
