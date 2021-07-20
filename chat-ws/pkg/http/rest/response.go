package rest

import (
	"encoding/json"
	"net/http"
)

type ResponseModel struct {
	Success      bool
	Data         interface{}
	ErrorMessage string
}

func NewResponseModel(success bool, data interface{}, errorMsg string) ResponseModel {
	return ResponseModel{
		Success:      success,
		Data:         data,
		ErrorMessage: errorMsg,
	}
}

func AddResponseToResponseWritter(rw http.ResponseWriter, data interface{}, errMsg string) {
	var success = false
	if errMsg == "" {
		success = true
	}
	json.NewEncoder(rw).Encode(NewResponseModel(success, data, errMsg))
}
