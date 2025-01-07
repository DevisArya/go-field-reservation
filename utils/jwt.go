package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userId int, email, role string) (string, error) {

	claims := jwt.MapClaims{
		"exp":     jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"iss":     "go-field-reservation",
		"sub":     fmt.Sprintf("%d", userId),
		"iat":     jwt.NewNumericDate(time.Now()),
		"user_id": userId,
		"email":   email,
		"role":    role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("SECRET_KEY environment variable not set")
	}

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
