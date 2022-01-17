package websocket

import "encoding/json"

const (
	UserJoin  = "USER_JOIN"
	UserLeave = "USER_LEAVE"
)

type Message struct {
	MessageType string          `json:"type"`
	Payload     json.RawMessage `json:"payload"`
}
