package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	ws "github.com/kjunn2000/straper/chat-ws/pkg/websocket"
)

var Upgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func SetUpRoutes() {
	pool := OpenPoolServer()
	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		HandleUpgrade(pool, rw, r)
	})
	http.HandleFunc("/chat-history",FetchChatHistoryHandler)
}

func OpenPoolServer() *ws.Pool {
	pool := ws.NewPool()
	go pool.StartWSServer()
	return pool
}

func HandleUpgrade(pool *ws.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Cannot upgrade to websocket connection : %s", err)
	}
	log.Println("Successful created websocket connection.")

	c := ws.NewClient(conn, pool)
	pool.Register <- c
	c.ReadMsg()
}

func FetchChatHistoryHandler(w http.ResponseWriter, r *http.Request){
	
	

}
