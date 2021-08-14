package account

import (
	"context"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, params CreateUserParam) error
	GetUserByUserId(ctx context.Context, userId string) (User, error)
	UpdateUser(ctx context.Context, param UpdateUserParam) error
	DeleteUser(ctx context.Context, userId string) error
}

type service struct {
	log *zap.Logger
	ur  Repository
}

func NewService(log *zap.Logger, ur Repository) *service {
	return &service{
		log: log,
		ur:  ur,
	}
}

func (us *service) Register(ctx context.Context, params CreateUserParam) error {
	hashedPassword, err := BcrptHashPassword(params.Password)
	if err != nil {
		return err
	}
	params.Password = hashedPassword
	params.Role = "USER"
	params.CreatedDate = time.Now()
	err = us.ur.CreateUser(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (us *service) GetUserByUserId(ctx context.Context, userId string) (User, error) {
	return us.ur.GetUserByUserId(ctx, userId)
}

func (us *service) UpdateUser(ctx context.Context, params UpdateUserParam) error {
	return us.ur.UpdateUser(ctx, params)
}

func (us *service) DeleteUser(ctx context.Context, userId string) error {
	return us.ur.DeleteUser(ctx, userId)
}

func BcrptHashPassword(password string) (string, error) {
	hashedBytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytePassword), nil
}
