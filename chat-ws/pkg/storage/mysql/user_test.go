package mysql

import (
	"time"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage"
)

func generateRandomoUser() *account.User {
	return &account.User{
		Username:    storage.RandomUsername(),
		Password:    storage.RandomPassword(),
		Role:        "USER",
		Email:       storage.RandomEmail(),
		PhoneNo:     storage.RandomPhoneNumber(),
		CreatedDate: time.Now(),
	}
}
