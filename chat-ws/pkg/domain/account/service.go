package account

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/kjunn2000/straper/chat-ws/configs"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, params CreateUserParam) error
	GetUserByUserId(ctx context.Context, userId string) (UserDetail, error)
	GetAccountListByWorkspaceId(ctx context.Context, workspaceId string) ([]UserInfo, error)
	UpdateUser(ctx context.Context, param UpdateUserParam) error

	ResetAccountPassword(ctx context.Context, email string) error
	UpdateAccountPassword(ctx context.Context, params UpdatePasswordParam) error

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

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (us *service) SeedUserAccount() error {
	usernames := []string{"Oliver", "Ivan", "Ava", "Josh", "John", "Harper", "Alexander", "Lucas", "Benjamin", "James",
		"William", "Emma", "Olivia", "Noah", "Liam", "Chris", "Belinda", "Casper", "Justin", "KaiXuan"}
	for i := 0; i < 20; i++ {
		index := strconv.Itoa(i)
		param := CreateUserParam{
			Username: usernames[i],
			Password: usernames[i] + "Password!",
			Email:    usernames[i] + "@gmail.com",
			PhoneNo:  "010123120" + index,
		}
		us.Register(context.Background(), param)
	}
	return nil
}

func (us *service) Register(ctx context.Context, params CreateUserParam) error {
	hashedPassword, err := BcrptHashPassword(params.Password)
	if err != nil {
		return err
	}
	params.Password = hashedPassword
	params.Role = RoleUser
	params.Status = StatusVerifying
	// params.Status = StatusActive
	params.CreatedDate = time.Now()
	err = us.ur.CreateUser(ctx, params)
	if err != nil {
		return us.verigyUserFieldError(err)
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

func (us *service) GetAccountListByWorkspaceId(ctx context.Context, workspaceId string) ([]UserInfo, error) {
	return us.ur.GetUserInfoListByWorkspaceId(ctx, workspaceId)
}

func (us *service) UpdateUser(ctx context.Context, params UpdateUserParam) error {
	params.UpdatedDate = time.Now()
	userDetail, err := us.GetUserByUserId(ctx, params.UserId)
	if err != nil {
		return err
	}
	if err := us.ur.UpdateUser(ctx, params); err != nil {
		return us.verigyUserFieldError(err)
	}
	if params.Email != userDetail.Email {
		if err := us.ur.UpdateAccountStatus(ctx, params.UserId, StatusVerifying); err != nil {
			return err
		}
		if err := us.CreateAndSendVerifyEmailToken(ctx, userDetail); err != nil {
			return err
		}
	}
	return nil
}

func BcrptHashPassword(password string) (string, error) {
	hashedBytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytePassword), nil
}

func (us *service) verigyUserFieldError(err error) error {
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
