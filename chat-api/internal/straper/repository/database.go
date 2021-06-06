package repository

import (
	"github.com/google/uuid"
	proto "github.com/kjunn2000/chat-app/chat-api/proto"
)

type MessageCrudRepository interface {
	CreateMessage(e *proto.Message) error
	ReadAllMessage() ([]*proto.Message, error)
	ReadMessage(id uuid.UUID) (*proto.Message, error)
	UpdateMessage(e *proto.Message) error
	DeleteMessage(id uuid.UUID) error
}
