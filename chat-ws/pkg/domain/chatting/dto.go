package chatting

type DeleteChatMessageParams struct {
	MessageId string `json:"message_id"`
	Type      string `json:"type"`
	Fid       string `json:"fid"`
}
