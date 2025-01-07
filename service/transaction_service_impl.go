package service

import (
	"context"
	"errors"

	"github.com/DevisArya/reservasi_lapangan/dto"
	"github.com/DevisArya/reservasi_lapangan/helper"
	"github.com/DevisArya/reservasi_lapangan/models"
	"github.com/DevisArya/reservasi_lapangan/repository"
	"github.com/DevisArya/reservasi_lapangan/utils"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	DB                    *gorm.DB
	validate              *validator.Validate
}

func NewTransactionService(transactionRepository repository.TransactionRepository, DB *gorm.DB, validate *validator.Validate) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		DB:                    DB,
		validate:              validate,
	}
}

// Save implements TransactionService
func (service *TransactionServiceImpl) Save(ctx context.Context, req *dto.TransactionReq, userId uint) (*dto.TransactionCreateResponse, error) {

	if err := service.validate.Struct(req); err != nil {
		return nil, err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// prepare a new transaction object
	newTransaction := models.Transaction{
		UserId: userId,
	}

	var total int64

	for _, detail := range req.TransactionDetail {
		newTransaction.TransactionDetail = append(newTransaction.TransactionDetail, models.TransactionDetail{
			ScheduleId: detail.ScheduleId,
			Price:      detail.Price,
		})

		total += detail.Price
	}

	newTransaction.TotalPrice = total

	//create midtrans payment url
	paymentUrl, err := utils.CreateMidtransUrl(&newTransaction)
	if err != nil {
		return nil, err
	}
	newTransaction.PaymentUrl = paymentUrl

	// save a new transaction data
	createdData, err := service.TransactionRepository.Save(ctx, tx, &newTransaction)
	if err != nil {
		return nil, err
	}

	// prepare response for the API
	res := dto.TransactionCreateResponse{
		Id:            createdData.Id,
		PaymentStatus: createdData.PaymentStatus,
		PaymentUrl:    createdData.PaymentUrl,
	}

	return &res, nil
}

// Update implements TransactionService
func (service *TransactionServiceImpl) Update(ctx context.Context, req *dto.MidtransRequest) error {
	if err := service.validate.Struct(req); err != nil {
		return err
	}

	Key, err := utils.Hash512(req.OrderId, req.StatusCode, req.GrossAmount)
	if err != nil {
		return err
	}

	if Key != req.SignatureKey {
		return errors.New("invalid transaction")
	}

	var status string

	if req.FraudStatus == "deny" {
		status = "Rejected"
	} else {
		if req.TransactionStatus == "capture" || req.TransactionStatus == "settlement" {
			status = "Success"
		} else if req.TransactionStatus == "deny" ||
			req.TransactionStatus == "cancel" ||
			req.TransactionStatus == "expire" ||
			req.TransactionStatus == "failure" {
			status = "Fail"
		} else if req.TransactionStatus == "pending" {
			status = "Pending"
		}
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	updatePayment := models.Transaction{
		OrderId:         req.OrderId,
		TransactionTime: req.TransactionTime,
		PaymentStatus:   status,
		PaymentType:     req.PaymentType,
		SettlementTime:  req.SettlementTime,
		FraudStatus:     req.FraudStatus,
	}

	if err := service.TransactionRepository.Update(ctx, tx, &updatePayment); err != nil {
		return err
	}

	return nil
}
