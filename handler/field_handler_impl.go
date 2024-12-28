package handler

import (
	"net/http"
	"strconv"

	"github.com/DevisArya/reservasi_lapangan/dto"
	"github.com/DevisArya/reservasi_lapangan/helper"
	"github.com/DevisArya/reservasi_lapangan/models"
	"github.com/DevisArya/reservasi_lapangan/service"
	"github.com/labstack/echo/v4"
)

type FieldHandlerImpl struct {
	FieldService service.FieldService
}

func NewFieldHandler(FieldService service.FieldService) FieldHandler {
	return &FieldHandlerImpl{
		FieldService,
	}
}

// Create implements FieldHandler
func (handler *FieldHandlerImpl) Create(c echo.Context) error {
	var req dto.FieldReqRes

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.NewResponse(http.StatusBadRequest, "Invalid request payload", nil))
	}

	field, err := handler.FieldService.Save(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.NewResponse(http.StatusInternalServerError, "Failed to create field", nil))
	}

	return c.JSON(
		http.StatusCreated,
		helper.NewResponse(http.StatusCreated, "Field created succesfully", field))
}

// Update implements FieldHandler
func (handler *FieldHandlerImpl) Update(c echo.Context) error {

	var req models.Field

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.NewResponse(http.StatusBadRequest, "Invalid field id", nil))
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.NewResponse(http.StatusBadRequest, "Invalid request payload", nil))
	}

	req.Id = uint(id)

	field, err := handler.FieldService.Update(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.NewResponse(http.StatusInternalServerError, "Failed to update field", nil))
	}

	return c.JSON(
		http.StatusOK,
		helper.NewResponse(http.StatusOK, "Field update succesfully", field))
}

// Delete implements FieldHandler
func (handler *FieldHandlerImpl) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.NewResponse(http.StatusBadRequest, "invalid id", nil))
	}

	if err := handler.FieldService.Delete(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.NewResponse(http.StatusInternalServerError, "Failed to delete field", nil))
	}

	return c.JSON(
		http.StatusOK,
		helper.NewResponse(http.StatusOK, "Field deleted sucessfully", nil))
}

// FindById implements FieldHandler
func (handler *FieldHandlerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.NewResponse(http.StatusBadRequest, "invalid id", nil))
	}

	field, err := handler.FieldService.FindById(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.NewResponse(http.StatusInternalServerError, "Failed to find field", nil))
	}

	return c.JSON(
		http.StatusOK,
		helper.NewResponse(http.StatusOK, "Success to find field", field))
}

// FindAll implements FieldHandler
func (handler *FieldHandlerImpl) FindAll(c echo.Context) error {
	fields, err := handler.FieldService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.NewResponse(http.StatusInternalServerError, "Failed to find fields", nil))
	}

	return c.JSON(
		http.StatusOK,
		helper.NewResponse(http.StatusOK, "Success to find fields", fields))
}
