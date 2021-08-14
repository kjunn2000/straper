package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) account.CreateUserParam {
	hashedPassword, err := account.BcrptHashPassword(storage.RandomPassword())
	require.NoError(t, err)

	createUserParams := account.CreateUserParam{
		Username:    storage.RandomUsername(),
		Password:    hashedPassword,
		Role:        "USER",
		Email:       storage.RandomEmail(),
		PhoneNo:     storage.RandomPhoneNumber(),
		CreatedDate: time.Now(),
	}

	err = store.CreateUser(context.Background(), createUserParams)
	require.NoError(t, err)
	return createUserParams
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByUsername(t *testing.T) {
	createUserParams := createRandomUser(t)
	user, err := store.GetUserByUsername(context.Background(), createUserParams.Username)
	require.NoError(t, err)
	require.Equal(t, createUserParams.Username, user.Username)
	require.Equal(t, createUserParams.Password, user.Password)
	require.Equal(t, createUserParams.Role, user.Role)
	require.Equal(t, createUserParams.Email, user.Email)
	require.Equal(t, createUserParams.PhoneNo, user.PhoneNo)

	require.NotZero(t, user.UserId)
}

func TestGetUserByUserId(t *testing.T) {
	createUserParams := createRandomUser(t)
	u, err := store.GetUserByUsername(context.Background(), createUserParams.Username)
	require.NoError(t, err)
	user, err := store.GetUserByUserId(context.Background(), u.UserId)
	require.NoError(t, err)
	require.Equal(t, createUserParams.Username, user.Username)
	require.Equal(t, createUserParams.Password, user.Password)
	require.Equal(t, createUserParams.Role, user.Role)
	require.Equal(t, createUserParams.Email, user.Email)
	require.Equal(t, createUserParams.PhoneNo, user.PhoneNo)

	require.NotZero(t, user.UserId)
	require.NotZero(t, user.CreatedDate)
}

func TestUpdateUser(t *testing.T) {
	createUserParams := createRandomUser(t)
	u, err := store.GetUserByUsername(context.Background(), createUserParams.Username)
	require.NoError(t, err)
	params := account.UpdateUserParam{
		UserId:   u.UserId,
		Username: storage.RandomString(8),
		Email:    storage.RandomEmail(),
		PhoneNo:  storage.RandomPhoneNumber(),
	}
	err = store.UpdateUser(context.Background(), params)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	createUserParams := createRandomUser(t)
	u, err := store.GetUserByUsername(context.Background(), createUserParams.Username)
	require.NoError(t, err)
	err = store.DeleteUser(context.Background(), u.UserId)
	require.NoError(t, err)
}
