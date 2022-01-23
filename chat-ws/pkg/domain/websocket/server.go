package websocket

type WSServer struct {
	activeUser map[string]*User
	register   chan *User
	unregister chan *User
	broadcast  chan *Message
}

func NewWSServer() *WSServer {
	return &WSServer{
		activeUser: make(map[string]*User),
		register:   make(chan *User),
		unregister: make(chan *User),
		broadcast:  make(chan *Message),
	}

}
