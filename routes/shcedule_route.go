package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	mdlw "github.com/DevisArya/reservasi_lapangan/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterScheduleRoutes(e *echo.Echo, scheduleHandler handler.ScheduleHandler) {

	// route for schedule
	e.POST("/schedule", scheduleHandler.Create, mdlw.Auth([]string{"admin", "user"}, false))
	e.PATCH("/schedule/:id", scheduleHandler.Update, mdlw.Auth([]string{"admin", "user"}, false))
	e.DELETE("/schedule/:id", scheduleHandler.Delete, mdlw.Auth([]string{"admin", "user"}, false))
	e.GET("/schedule/:id", scheduleHandler.FindById)
	e.GET("/schedules", scheduleHandler.FindAll)
}
