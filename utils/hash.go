package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"os"
)

func Hash512(orderId, statusCode, grossAmount string) (string, error) {

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		return "", errors.New("env MIDTRANS_SERVER_KEY is not set")
	}

	input := orderId + statusCode + grossAmount + serverKey
	inputBytes := []byte(input)
	sha512Hasher := sha512.New()
	sha512Hasher.Write(inputBytes)
	hashedInputBytes := sha512Hasher.Sum(nil)
	hashedInputString := hex.EncodeToString(hashedInputBytes)

	return hashedInputString, nil
}
