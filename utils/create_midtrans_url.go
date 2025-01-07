package utils

import (
	"errors"
	"os"
	"strconv"

	"github.com/DevisArya/reservasi_lapangan/models"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func CreateMidtransUrl(transaction *models.Transaction) (string, error) {

	var s = snap.Client{}
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		return "", errors.New("env MIDTRANS_SERVER_KEY is not set")
	}
	s.New(serverKey, midtrans.Sandbox)

	var items []midtrans.ItemDetails

	for _, val := range transaction.TransactionDetail {
		item := midtrans.ItemDetails{
			Name:  val.Name,
			Price: val.Price,
		}
		items = append(items, item)
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.FormatUint(uint64(transaction.Id), 10),
			GrossAmt: int64(transaction.TotalPrice),
		},
		Items: &items,
	}

	snapRes, err := s.CreateTransactionUrl(req)
	if err != nil {
		return "", err
	}

	return snapRes, nil
}
