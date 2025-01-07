package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/DevisArya/reservasi_lapangan/dto"
	"github.com/DevisArya/reservasi_lapangan/helper"
	"github.com/DevisArya/reservasi_lapangan/repository"
	"github.com/DevisArya/reservasi_lapangan/utils"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *gorm.DB
	validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, DB *gorm.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             DB,
		validate:       validate,
	}
}

// Login implements AuthService
func (service *AuthServiceImpl) Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error) {

	fmt.Println(request)
	if err := service.validate.Struct(request); err != nil {
		return nil, err
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.AuthRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		return nil, err
	}

	if !utils.ComparePassword(user.Password, request.Password) {
		return nil, errors.New("wrong password")
	}

	//create jwt token
	token, err := utils.CreateToken(int(user.Id), user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	res := &dto.LoginResponse{
		User: dto.UserLoginResponse{
			Id:       user.Id,
			Username: user.Name,
			Role:     user.Role,
		},
		Token: token,
	}

	return res, nil
}
