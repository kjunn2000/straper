package websocket

import "encoding/json"

const (
	UserJoin  = "USER_JOIN"
	UserLeave = "USER_LEAVE"
)

type Message struct {
	MessageType string          `json:"message_type"`
	WorkspaceId string          `json:"workspace_id"`
	ChannelId   string          `json:"channel_id"`
	SenderId    string          `json:"sender_id"`
	UserDetail  UserDetail      `json:"user_detail"`
	Payload     json.RawMessage `json:"payload"`
}

type UserDetail struct {
	UserId   string `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email" validate:"email"`
	PhoneNo  string `json:"phone_no" db:"phone_no"`
}
