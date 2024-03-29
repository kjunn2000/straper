package chatting

import (
	"encoding/json"
	"time"
)

type Message struct {
	MessageId   string     `json:"message_id" db:"message_id"`
	Type        string     `json:"type" db:"type"`
	ChannelId   string     `json:"channel_id" db:"channel_id"`
	CreatorId   string     `json:"creator_id" db:"creator_id"`
	UserDetail  UserDetail `json:"user_detail"`
	Content     string     `json:"content" db:"content"`
	FileName    string     `json:"file_name" db:"file_name"`
	FileType    string     `json:"file_type" db:"file_type"`
	FileBytes   []byte     `json:"file_bytes"`
	CreatedDate time.Time  `json:"created_date" db:"created_date"`
	Cursor      string     `json:"cursor"`
}

type UserDetail struct {
	UserId   string `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email" validate:"email"`
	PhoneNo  string `json:"phone_no" db:"phone_no"`
}

func (message *Message) Encode() ([]byte, error) {
	json, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	return json, nil
}
