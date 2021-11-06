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
		Status:      "VERIFYING",
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

func TestGetUserDetailByUsername(t *testing.T) {
	createUserParams := createRandomUser(t)
	user, err := store.GetUserDetailByUsername(context.Background(), createUserParams.Username)
	require.NoError(t, err)
	require.Equal(t, createUserParams.Username, user.Username)
	require.Equal(t, createUserParams.Email, user.Email)
	require.Equal(t, createUserParams.PhoneNo, user.PhoneNo)
	require.Equal(t, createUserParams.CreatedDate.Day(), user.CreatedDate.Day())
	require.Equal(t, createUserParams.CreatedDate.Hour(), user.CreatedDate.Hour())
	require.Equal(t, createUserParams.CreatedDate.Minute(), user.CreatedDate.Minute())

	require.NotZero(t, user.UserId)
}

func TestGetUserDetailByUserId(t *testing.T) {
	createUserParams := createRandomUser(t)
	user, _ := store.GetUserDetailByUsername(context.Background(), createUserParams.Username)
	user, err := store.GetUserDetailByUserId(context.Background(), user.UserId)
	require.NoError(t, err)
	require.Equal(t, createUserParams.Username, user.Username)
	require.Equal(t, createUserParams.Email, user.Email)
	require.Equal(t, createUserParams.PhoneNo, user.PhoneNo)
	require.Equal(t, createUserParams.CreatedDate.Day(), user.CreatedDate.Day())
	require.Equal(t, createUserParams.CreatedDate.Hour(), user.CreatedDate.Hour())
	require.Equal(t, createUserParams.CreatedDate.Minute(), user.CreatedDate.Minute())

	require.NotZero(t, user.UserId)
}

func TestGetUserCredentialByUsername(t *testing.T) {
	createUserParams := createRandomUser(t)
	user, err := store.GetUserCredentialByUsername(context.Background(), createUserParams.Username)
	require.NoError(t, err)
	require.Equal(t, createUserParams.Username, user.Username)
	require.Equal(t, createUserParams.Password, user.Password)
	require.Equal(t, createUserParams.Status, user.Status)
	require.Equal(t, createUserParams.CreatedDate.Day(), user.CreatedDate.Day())
	require.Equal(t, createUserParams.CreatedDate.Hour(), user.CreatedDate.Hour())
	require.Equal(t, createUserParams.CreatedDate.Minute(), user.CreatedDate.Minute())

	require.NotZero(t, user.UserId)
}

func TestGetUserCredentialByUserId(t *testing.T) {
	createUserParams := createRandomUser(t)
	u, _ := store.GetUserDetailByUsername(context.Background(), createUserParams.Username)
	user, err := store.GetUserCredentialByUserId(context.Background(), u.UserId)
	require.NoError(t, err)
	require.Equal(t, createUserParams.Password, user.Password)
	require.Equal(t, createUserParams.Status, user.Status)
	require.Equal(t, createUserParams.CreatedDate.Day(), user.CreatedDate.Day())
	require.Equal(t, createUserParams.CreatedDate.Hour(), user.CreatedDate.Hour())
	require.Equal(t, createUserParams.CreatedDate.Minute(), user.CreatedDate.Minute())

	require.NotZero(t, user.UserId)
}

func TestUpdateUser(t *testing.T) {
	createUserParams := createRandomUser(t)
	user, _ := store.GetUserDetailByUsername(context.Background(), createUserParams.Username)
	params := account.UpdateUserParam{
		UserId:      user.UserId,
		Username:    storage.RandomString(8),
		Email:       storage.RandomEmail(),
		PhoneNo:     storage.RandomPhoneNumber(),
		UpdatedDate: time.Now(),
	}
	err := store.UpdateUser(context.Background(), params)
	require.NoError(t, err)
}

func TestUpdateAccountStatus(t *testing.T) {
	createUserParams := createRandomUser(t)
	u, err := store.GetUserDetailByUsername(context.Background(), createUserParams.Username)
	require.NoError(t, err)
	err = store.UpdateAccountStatus(context.Background(), u.UserId, "ACTIVE")
	require.NoError(t, err)
}

func TestUpdateAccountPassword(t *testing.T) {
	createUserParams := createRandomUser(t)
	u, err := store.GetUserDetailByUsername(context.Background(), createUserParams.Username)
	require.NoError(t, err)
	hashedPassword, err := account.BcrptHashPassword(storage.RandomPassword())
	require.NoError(t, err)
	err = store.UpdateAccountPassword(context.Background(), u.UserId, hashedPassword)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	createUserParams := createRandomUser(t)
	u, err := store.GetUserDetailByUsername(context.Background(), createUserParams.Username)
	require.NoError(t, err)
	err = store.DeleteUser(context.Background(), u.UserId)
	require.NoError(t, err)
}
