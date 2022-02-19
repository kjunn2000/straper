package chatting

type EditChatMessageParams struct {
	MessageId string `json:"message_id"`
	Content   string `json:"content" db:"content"`
}

type DeleteChatMessageParams struct {
	MessageId string `json:"message_id"`
	Type      string `json:"type"`
	Fid       string `json:"fid"`
}
