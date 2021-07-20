package account

import (
	"database/sql"
	"time"

	"go.uber.org/zap"
)

type Repository interface {
	SaveUser(user *User) error
	CheckUsernameExist(username string) (bool, error)
}

type Service interface {
	Register(user User) error
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

func (us *service) Register(user User) error {
	exist, err := us.ur.CheckUsernameExist(user.Username)

	if err != sql.ErrNoRows || exist {
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
