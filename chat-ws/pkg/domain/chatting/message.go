package chatting

import (
	"encoding/json"
	"time"
)

const (
	UserJoin  = "USER_JOIN"
	UserLeave = "USER_LEAVE"
	Messaging = "MESSAGE"
	File      = "FILE"
)

type Message struct {
	MessageId   string    `json:"message_id" db:"message_id"`
	Type        string    `json:"type" db:"type"`
	ChannelId   string    `json:"channel_id" db:"channel_id"`
	CreatorName string    `json:"creator_name" db:"creator_name"`
	Content     string    `json:"content" db:"content"`
	FileName    string    `json:"file_name" db:"file_name"`
	FileType    string    `json:"file_type" db:"file_type"`
	FileBytes   []byte    `json:"file_bytes"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
}

func (message *Message) Encode() ([]byte, error) {
	json, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	return json, nil
}
