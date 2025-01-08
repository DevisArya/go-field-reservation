package service

import (
	"context"
	"errors"
	"strconv"
	"time"

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

	transId := "eco" + strconv.FormatUint(uint64(userId), 10) + time.Now().UTC().Format("2006010215040105")
	// prepare a new transaction object
	newTransaction := models.Transaction{
		PaymentStatus:   "unpaid",
		TransactionId:   transId,
		UserId:          userId,
		TransactionTime: time.Now(),
	}

	var total int64

	for _, detail := range req.TransactionDetail {
		newTransaction.TransactionDetail = append(newTransaction.TransactionDetail, models.TransactionDetail{
			ScheduleId: detail.ScheduleId,
			Price:      detail.Price,
			Name:       detail.Name,
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
		TransactionId: createdData.TransactionId,
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

	Key, err := utils.Hash512(req.TransactionId, req.StatusCode, req.GrossAmount)
	if err != nil {
		return err
	}

	if Key != req.SignatureKey {
		return errors.New("invalid transaction")
	}

	var status string

	if req.FraudStatus == "deny" {
		status = "rejected"
	} else {
		if req.TransactionStatus == "capture" || req.TransactionStatus == "settlement" {
			status = "success"
		} else if req.TransactionStatus == "deny" ||
			req.TransactionStatus == "cancel" ||
			req.TransactionStatus == "expire" ||
			req.TransactionStatus == "failure" {
			status = "fail"
		} else if req.TransactionStatus == "pending" {
			status = "pending"
		}
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	layout := "2006-01-02 15:04:05"

	// Parsing string ke time.Time
	transTime, err := time.Parse(layout, req.TransactionTime)
	if err != nil {
		return err
	}
	// Parsing string ke time.Time
	settlementTime, err := time.Parse(layout, req.SettlementTime)
	if err != nil {
		return err
	}

	updatePayment := models.Transaction{
		TransactionId:   req.TransactionId,
		OrderId:         req.OrderId,
		TransactionTime: transTime,
		PaymentStatus:   status,
		PaymentType:     req.PaymentType,
		SettlementTime:  settlementTime,
		FraudStatus:     req.FraudStatus,
	}

	if err := service.TransactionRepository.Update(ctx, tx, &updatePayment); err != nil {
		return err
	}

	return nil
}
