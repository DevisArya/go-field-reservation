package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetIdFromClaims(c echo.Context) (uint, error) {
	// Ambil claims dari context
	claimsInterface := c.Get("claims")
	if claimsInterface == nil {
		return 0, errors.New("claims not found in context")
	}

	// Type assertion langsung ke jwt.MapClaims
	userClaims, ok := claimsInterface.(*jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims format")
	}

	// Ambil user_id dan konversi ke uint
	userIdFloat, ok := (*userClaims)["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id not found or invalid type in claims")
	}

	// Konversi dari float64 ke uint
	userId := uint(userIdFloat)

	return userId, nil
}
