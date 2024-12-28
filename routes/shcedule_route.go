package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	"github.com/labstack/echo/v4"
)

func RegisterScheduleRoutes(e *echo.Echo, scheduleHandler handler.ScheduleHandler) {

	// route for schedule
	e.POST("/schedule", scheduleHandler.Create)
	e.PATCH("/schedule/:id", scheduleHandler.Update)
	e.DELETE("/schedule/:id", scheduleHandler.Delete)
	e.GET("/schedule/:id", scheduleHandler.FindById)
	e.GET("/schedules", scheduleHandler.FindAll)
}
