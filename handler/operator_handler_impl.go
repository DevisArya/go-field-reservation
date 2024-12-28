package handler

import (
	"net/http"
	"strconv"

	"github.com/DevisArya/reservasi_lapangan/dto"
	"github.com/DevisArya/reservasi_lapangan/models"
	"github.com/DevisArya/reservasi_lapangan/service"
	"github.com/labstack/echo/v4"
)

type OperatorHandlerImpl struct {
	OperatorService service.OperatorService
}

func NewOperatorHandler(operatorService service.OperatorService) OperatorHandler {
	return &OperatorHandlerImpl{
		OperatorService: operatorService,
	}
}

// Create implements OperatorHandler
func (handler *OperatorHandlerImpl) Create(c echo.Context) error {
	var req dto.OperatorCreateRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := handler.OperatorService.Create(c.Request().Context(), &req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "Operator created sucessfully")
}

// Update implements OperatorHandler
func (handler *OperatorHandlerImpl) Update(c echo.Context) error {
	var req dto.OperatorUpdateRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid operator ID")
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	var updateRequest = models.Operator{
		Id:       uint(id),
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	if err := handler.OperatorService.Update(c.Request().Context(), &updateRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Operator update sucessfully")
}

// Delete implements OperatorHandler
func (handler *OperatorHandlerImpl) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid operator ID")
	}

	if err := handler.OperatorService.Delete(c.Request().Context(), uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Operator Deleted sucessfully")
}

// FindById implements OperatorHandler
func (handler *OperatorHandlerImpl) FindById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid operator ID")
	}

	operator, err := handler.OperatorService.FindById(c.Request().Context(), uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, operator)
}

// FindAll implements OperatorHandler
func (handler *OperatorHandlerImpl) FindAll(c echo.Context) error {

	operators, err := handler.OperatorService.FindAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, operators)
}
