package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	"github.com/labstack/echo/v4"
)

func RegisterTransactionRoutes(e *echo.Echo, handler handler.TransactionHandler) {

	// route for transaction
	e.POST("/transaction", handler.Create)
	e.PATCH("/transaction/update", handler.Create)
}
