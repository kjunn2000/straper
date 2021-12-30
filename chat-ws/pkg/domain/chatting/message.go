package chatting

import "encoding/json"

const (
	UserJoin  = "UserJoin"
	UserLeave = "UserLeave"
	Messaging = "Message"
)

type Message struct {
	Type      string `json:"type"`
	ChannelId string `json:"channel_id"`
	Content   string `json:"content"`
}

func (message *Message) Encode() ([]byte, error) {
	json, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	return json, nil
}
