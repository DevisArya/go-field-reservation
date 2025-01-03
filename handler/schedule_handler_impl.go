package handler

import (
	"net/http"
	"strconv"

	"github.com/DevisArya/reservasi_lapangan/models"
	"github.com/DevisArya/reservasi_lapangan/service"
	"github.com/labstack/echo/v4"
)

type ScheduleHandlerImpl struct {
	ScheduleService service.ScheduleService
}

func NewScheduleHandler(scheduleService service.ScheduleService) ScheduleHandler {
	return &ScheduleHandlerImpl{
		ScheduleService: scheduleService,
	}
}

// Create implements ScheduleHandler
func (handler *ScheduleHandlerImpl) Create(c echo.Context) error {
	var req models.Schedule

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := handler.ScheduleService.Create(c.Request().Context(), &req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "Schedule created sucessfully")
}

// Update implements ScheduleHandler
func (handler *ScheduleHandlerImpl) Update(c echo.Context) error {
	var req models.Schedule

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid schedule ID")
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	req.Id = uint(id)

	if err := handler.ScheduleService.Update(c.Request().Context(), &req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Schedule update sucessfully")
}

// Delete implements ScheduleHandler
func (handler *ScheduleHandlerImpl) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid schedule ID")
	}

	if err := handler.ScheduleService.Delete(c.Request().Context(), uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Schedule Deleted sucessfully")
}

// FindById implements ScheduleHandler
func (handler *ScheduleHandlerImpl) FindById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid schedule ID")
	}

	schedule, err := handler.ScheduleService.FindById(c.Request().Context(), uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schedule)
}

// FindAll implements ScheduleHandler
func (handler *ScheduleHandlerImpl) FindAll(c echo.Context) error {

	schedules, err := handler.ScheduleService.FindAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, schedules)
}
