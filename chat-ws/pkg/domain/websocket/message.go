package websocket

import "encoding/json"

type Message struct {
	MessageType string          `json:"type"`
	Payload     json.RawMessage `json:"payload"`
	WorkspaceId string          `json:"workspace_id"`
	SenderId    string          `json:"sender_id"`
}
