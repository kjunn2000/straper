package main

import (
	"net/http"

	handler "github.com/kjunn2000/straper/chat-ws/pkg/http"
)

func main() {
	handler.SetUpRoutes()
	http.ListenAndServe(":9090", nil)
}
