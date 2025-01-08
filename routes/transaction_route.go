package routes

import (
	"github.com/DevisArya/reservasi_lapangan/handler"
	mdlw "github.com/DevisArya/reservasi_lapangan/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterTransactionRoutes(e *echo.Echo, handler handler.TransactionHandler) {

	// route for transaction
	e.POST("/transaction", handler.Create, mdlw.Auth([]string{"user"}, false))
	e.POST("/transaction/update", handler.Update, mdlw.Auth([]string{"user"}, false))
}
