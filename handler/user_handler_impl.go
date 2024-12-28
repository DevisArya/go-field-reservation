package handler

import (
	"net/http"
	"strconv"

	"github.com/DevisArya/reservasi_lapangan/dto"
	"github.com/DevisArya/reservasi_lapangan/models"
	"github.com/DevisArya/reservasi_lapangan/service"
	"github.com/labstack/echo/v4"
)

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &UserHandlerImpl{
		UserService: userService,
	}
}

// Create implements UserHandler
func (handler *UserHandlerImpl) Create(c echo.Context) error {
	var req dto.UserCreateRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := handler.UserService.Create(c.Request().Context(), &req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "User created sucessfully")
}

// Update implements UserHandler
func (handler *UserHandlerImpl) Update(c echo.Context) error {
	var req dto.UserUpdateRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	var updateRequest = models.User{
		Id:       uint(id),
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	if err := handler.UserService.Update(c.Request().Context(), &updateRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User update sucessfully")
}

// Delete implements UserHandler
func (handler *UserHandlerImpl) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	if err := handler.UserService.Delete(c.Request().Context(), uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User Deleted sucessfully")
}

// FindById implements UserHandler
func (handler *UserHandlerImpl) FindById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := handler.UserService.FindById(c.Request().Context(), uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// FindAll implements UserHandler
func (handler *UserHandlerImpl) FindAll(c echo.Context) error {

	users, err := handler.UserService.FindAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
