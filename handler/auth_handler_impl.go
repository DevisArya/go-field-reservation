package handler

import (
	"net/http"

	"github.com/DevisArya/reservasi_lapangan/dto"
	"github.com/DevisArya/reservasi_lapangan/helper"
	"github.com/DevisArya/reservasi_lapangan/service"
	"github.com/labstack/echo/v4"
)

type AuthHandlerImpl struct {
	AuthService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &AuthHandlerImpl{
		AuthService: authService,
	}
}

// Login implements AuthHandler
func (handler *AuthHandlerImpl) Login(c echo.Context) error {

	var request dto.LoginRequest
	err := c.Bind(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helper.NewResponse(http.StatusBadRequest, "Invalid Payload", nil))
	}

	res, err := handler.AuthService.Login(c.Request().Context(), &request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.NewResponse(http.StatusInternalServerError, err.Error(), nil))
		// helper.NewResponse(http.StatusInternalServerError, "Login Failed", nil))
	}
	// }

	return c.JSON(http.StatusOK,
		helper.NewResponse(http.StatusOK, "Login Succesfully", res))
}
