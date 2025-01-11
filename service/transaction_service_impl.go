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

	// validate struct request
	if err := service.validate.Struct(req); err != nil {
		return nil, err
	}

	// transactional database
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	//generate transaction id
	transId := "eco" + strconv.FormatUint(uint64(userId), 10) + time.Now().UTC().Format("2006010215040105")

	// prepare a new transaction object
	newTransaction := models.Transaction{
		PaymentStatus:   "unpaid",
		TransactionId:   transId,
		UserId:          userId,
		TransactionTime: time.Now(),
	}

	var total int64
	scheduleIds := make([]uint, len(req.TransactionDetail))

	for i, detail := range req.TransactionDetail {
		newTransaction.TransactionDetail = append(newTransaction.TransactionDetail, models.TransactionDetail{
			ScheduleId: detail.ScheduleId,
			Price:      detail.Price,
			Name:       detail.Name,
		})

		total += detail.Price
		scheduleIds[i] = detail.ScheduleId
	}
	newTransaction.TotalPrice = total

	// check status schedule (available or not)
	lenSchedule, err := service.TransactionRepository.GetScheduleByIds(ctx, tx, scheduleIds)
	if err != nil {
		return nil, err
	} else if lenSchedule != len(scheduleIds) {
		return nil, errors.New("one or more schedules are not available")
	}

	// lock and update status schedule
	if err := service.TransactionRepository.LockScheduleAndUpdate(ctx, tx, scheduleIds); err != nil {
		return nil, err
	}

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
	// Validasi request
	if err := service.validate.Struct(req); err != nil {
		return err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	// Generate hash key dan validasi signature
	Key, err := utils.Hash512(req.TransactionId, req.StatusCode, req.GrossAmount)
	if err != nil {
		return err
	}

	if Key != req.SignatureKey {
		return errors.New("invalid transaction")
	}

	// Ambil data transaksi dari database
	transaction, err := service.TransactionRepository.GetTransactionById(ctx, tx, req.TransactionId)
	if err != nil {
		return err
	}

	scheduleIds := make([]uint, len(transaction.TransactionDetail))
	for i, detail := range transaction.TransactionDetail {
		scheduleIds[i] = detail.ScheduleId
	}

	// Tentukan status pembayaran dan jadwal berdasarkan status transaksi
	var (
		statusPayment  string
		statusSchedule string
	)

	if req.FraudStatus == "deny" {
		statusPayment, statusSchedule = "rejected", "available"
	} else {
		switch req.TransactionStatus {
		case "capture", "settlement":
			statusPayment, statusSchedule = "success", "sold"

		case "deny", "cancel", "expire", "failure":
			statusPayment, statusSchedule = "rejected", "available"

		case "pending":
			statusPayment = "pending"
		}
	}

	// Parsing waktu transaksi dan waktu settlement
	layout := "2006-01-02 15:04:05"
	transTime, err := time.Parse(layout, req.TransactionTime)
	if err != nil {
		return err
	}

	settlementTime, err := time.Parse(layout, req.SettlementTime)
	if err != nil {
		return err
	}

	// Update data pembayaran
	updatePayment := models.Transaction{
		TransactionId:   req.TransactionId,
		OrderId:         req.OrderId,
		TransactionTime: transTime,
		PaymentStatus:   statusPayment,
		PaymentType:     req.PaymentType,
		SettlementTime:  settlementTime,
		FraudStatus:     req.FraudStatus,
	}

	if err := service.TransactionRepository.Update(ctx, tx, &updatePayment); err != nil {
		return err
	}

	// Update status jadwal jika pembayaran tidak dalam status pending
	if statusPayment != "pending" {
		if err := service.TransactionRepository.UpdateSchedulesStatus(ctx, tx, scheduleIds, statusSchedule); err != nil {
			return err
		}
	}

	return nil
}
