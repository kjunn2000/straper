package data

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kjunn2000/chat-app/chat-api/config"
	proto "github.com/kjunn2000/chat-app/chat-api/proto"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type MessageStore struct {
	Db  *sqlx.DB
	Log *zap.Logger
}

func NewMessageStore(logger *zap.Logger) *MessageStore {
	logger.Info("connstring", zap.String("string", config.NewConnString()))
	db, err := sqlx.Connect("postgres", config.NewConnString())
	if err != nil {
		logger.Fatal("Cannot establish connection to db", zap.Error(err))
	}
	logger.Info("Connected to db")
	return &MessageStore{
		Db:  db,
		Log: logger,
	}
}

func (s *MessageStore) CreateMessage(e *proto.Message) error {
	_, err := s.Db.Exec(
		`INSERT INTO messages (id,content,created_at) values($1, $2, $3);`,
		e.Id,
		e.Content,
		e.Timestamp)
	if err != nil {
		s.Log.Fatal("Failed to insert data to db", zap.Error(err))
	}
	fmt.Printf("Successful insert %s", e.Id)
	return nil
}

func (s *MessageStore) ReadAllMessage() ([]*proto.Message, error) {
	var ms []*proto.Message
	err := s.Db.Select(
		&ms,
		`SELECT * FROM messages;`,
	)
	if err != nil {
		log.Fatalln("Unable to select all from db")
	}
	return ms, nil
}

func (s *MessageStore) ReadMessage(id uuid.UUID) (*proto.Message, error) {
	var m proto.Message
	err := s.Db.Get(
		&m,
		`SELECT * FROM message where id = $1;`,
		id,
	)
	if err != nil {
		log.Fatalln("Unable to select one from db")
	}
	return &m, nil
}

func (s *MessageStore) UpdateMessage(e *proto.Message) error {
	_, err := s.Db.Exec(
		`UPDATE messages SET content = $1, timestamp = $2 WHERE id = $3;`,
	)
	if err != nil {
		log.Fatalln("Cannot udpate message to db")
	}
	return nil
}

func (s *MessageStore) DeleteMessage(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}
