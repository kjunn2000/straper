package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/kjunn2000/straper/chat-ws/configs"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestRegister(t *testing.T) {
	acc := getRandomAccount()

	testCases := []struct {
		name          string
		buildStubs    func(store *mock.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(store *mock.MockStore) {
				gomock.InOrder(
					store.EXPECT().
						CheckUsernameExist(acc.Username).
						Times(1).
						Return(false, sql.ErrNoRows),
					store.EXPECT().
						SaveUser(gomock.Any()).
						Times(1).
						Return(nil),
				)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockStore(ctrl)
			tc.buildStubs(store)

			log, _ := zap.NewDevelopment()
			config, err := configs.LoadConfig("../../../../")
			require.NoError(t, err)
			server, err := NewServer(log, config, store)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()

			url := "http://localhost:8080/api/v1/account/opening"
			requestBytes, _ := json.Marshal(&acc)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBytes))
			require.NoError(t, err)

			server.httpServer.Handler.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func getRandomAccount() *account.User {
	return &account.User{
		UserId:      storage.RandomUserId(),
		Username:    storage.RandomUsername(),
		Password:    storage.RandomPassword(),
		Role:        "USER",
		Email:       storage.RandomEmail(),
		PhoneNo:     storage.RandomPhoneNumber(),
		CreatedDate: time.Now(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, acc account.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var newAccount account.User
	err = json.Unmarshal(data, &newAccount)
	require.NoError(t, err)
	require.Equal(t, acc, newAccount)
}
