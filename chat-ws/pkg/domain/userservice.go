package domain

import (
	"database/sql"
	"time"

	"go.uber.org/zap"
)

type UserRepository interface {
	SaveUser(user *User) error
	FindUserByUsername(username string) (*User, error)
}

type UserService interface {
	Register(user User) error
}

type userService struct {
	log *zap.Logger
	ur  UserRepository
}

func NewUserService(log *zap.Logger, ur UserRepository) *userService {
	return &userService{
		log: log,
		ur:  ur,
	}

}

func (us *userService) Register(user User) error {
	_, err := us.ur.FindUserByUsername(user.Username)

	if err != sql.ErrNoRows {
		us.log.Info("Username already exist.", zap.Error(err))
		return err
	}

	user.Role = "USER"
	user.CreatedDate = time.Now()
	err = us.ur.SaveUser(&user)
	if err != nil {
		return err
	}
	return nil
}
