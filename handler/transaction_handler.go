package handler

import "github.com/labstack/echo/v4"

type TransactionHandler interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
}
