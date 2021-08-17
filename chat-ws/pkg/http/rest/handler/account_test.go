package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/kjunn2000/straper/chat-ws/configs"
	"github.com/kjunn2000/straper/chat-ws/pkg/domain/account"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage"
	"github.com/kjunn2000/straper/chat-ws/pkg/storage/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func EqUser(user account.CreateUserParam, password string) gomock.Matcher {
	return eqUserMatcher{user: user, password: password}
}

type eqUserMatcher struct {
	user     account.CreateUserParam
	password string
}

func (e eqUserMatcher) Matches(x interface{}) bool {
	arg, ok := x.(account.CreateUserParam)
	if !ok {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(arg.Password), []byte(e.password))
	if err != nil {
		return false
	}
	e.user.Password = arg.Password
	if e.user.CreatedDate.Round(time.Second).Equal(arg.CreatedDate.Round(time.Second)) {
		e.user.CreatedDate = arg.CreatedDate
	}
	return reflect.DeepEqual(e.user, arg)
}

func (e eqUserMatcher) String() string {
	return fmt.Sprintf("is equal to user : %v and password : %v", e.user, e.user.Password)
}

func getRandomAccount() account.CreateUserParam {
	return account.CreateUserParam{
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
				store.EXPECT().
					CreateUser(gomock.Any(), EqUser(acc, acc.Password)).
					Times(1).
					Return(nil)
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

			url := "http://localhost:8080/api/v1/account/create"
			requestBytes, _ := json.Marshal(&acc)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBytes))
			require.NoError(t, err)

			server.httpServer.Handler.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
