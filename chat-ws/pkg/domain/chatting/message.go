package chatting

import "encoding/json"

const (
	UserJoinedAction = "UserJoined"
	UserLeaveAction  = "UserLeave"
	MessageAction    = "Message"
)

type Message struct {
	Action    string `json:"action"`
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
