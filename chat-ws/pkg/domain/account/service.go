package account

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/kjunn2000/straper/chat-ws/configs"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, params CreateUserParam) error
	GetUserByUserId(ctx context.Context, userId string) (UserDetail, error)
	UpdateUser(ctx context.Context, param UpdateUserParam) error
	DeleteUser(ctx context.Context, userId string) error
	ValidateVerifyEmailToken(ctx context.Context, tokenId string) error
}

type service struct {
	log    *zap.Logger
	ur     Repository
	config configs.Config
}

func NewService(log *zap.Logger, ur Repository, config configs.Config) *service {
	return &service{
		log:    log,
		ur:     ur,
		config: config,
	}
}

func (us *service) Register(ctx context.Context, params CreateUserParam) error {
	hashedPassword, err := BcrptHashPassword(params.Password)
	if err != nil {
		return err
	}
	params.Password = hashedPassword
	params.Role = RoleUser
	params.Status = StatusVerifying
	params.CreatedDate = time.Now()
	err = us.ur.CreateUser(ctx, params)
	if err != nil {
		fetchField := strings.Split(err.Error(), ".")
		field := fetchField[len(fetchField)-1]
		field = field[:len(field)-1]
		switch field {
		case "username":
			return errors.New("username.registered")
		case "email":
			return errors.New("email.registered")
		case "phone_no":
			return errors.New("phone.no.registered")
		default:
			return err
		}
	}
	user, err := us.ur.GetUserDetailByUsername(ctx, params.Username)
	if err != nil {
		return err
	}
	err = us.CreateAndSendVerifyEmailToken(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (us *service) GetUserByUserId(ctx context.Context, userId string) (UserDetail, error) {
	return us.ur.GetUserDetailByUserId(ctx, userId)
}

func (us *service) UpdateUser(ctx context.Context, params UpdateUserParam) error {
	params.UpdatedDate = time.Now()
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
