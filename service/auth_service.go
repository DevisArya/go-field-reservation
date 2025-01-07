package service

import (
	"context"

	"github.com/DevisArya/reservasi_lapangan/dto"
)

type AuthService interface {
	Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error)
}
