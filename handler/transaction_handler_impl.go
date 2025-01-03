package handler

import (
	"net/http"

	"github.com/DevisArya/reservasi_lapangan/dto"
	"github.com/DevisArya/reservasi_lapangan/helper"
	"github.com/DevisArya/reservasi_lapangan/service"
	"github.com/labstack/echo/v4"
)

type TransactionHandlerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) TransactionHandler {
	return &TransactionHandlerImpl{
		TransactionService: transactionService,
	}
}

// Create implements TransactionHandler
func (handler *TransactionHandlerImpl) Create(c echo.Context) error {

	var transactionData dto.TransactionReq

	if err := c.Bind(&transactionData); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.NewResponse(http.StatusBadRequest, "Invalid request payload", nil))
	}

	res, err := handler.TransactionService.Save(c.Request().Context(), &transactionData)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.NewResponse(http.StatusInternalServerError, "Failed to create transaction", nil))
	}

	return c.JSON(
		http.StatusCreated,
		helper.NewResponse(http.StatusCreated, "Transaction created successfully", res))
}

// Update implements TransactionHandler
func (handler *TransactionHandlerImpl) Update(c echo.Context) error {

	var transactionData dto.MidtransRequest

	if err := c.Bind(&transactionData); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.NewResponse(http.StatusBadRequest, "Invalid request payload", nil))
	}

	err := handler.TransactionService.Update(c.Request().Context(), &transactionData)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.NewResponse(http.StatusInternalServerError, "Failed to update transaction", nil))
	}

	return c.JSON(
		http.StatusOK,
		helper.NewResponse(http.StatusOK, "Transaction update successfully", nil))
}
