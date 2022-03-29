package chatting

import "time"

type EditChatMessageParams struct {
	MessageId string `json:"message_id"`
	Content   string `json:"content" db:"content"`
}

type DeleteChatMessageParams struct {
	MessageId string `json:"message_id"`
	Type      string `json:"type"`
	Fid       string `json:"fid"`
}

type PaginationMessagesParam struct {
	Cursor      string
	CreatedTime time.Time
	Id          string
}
