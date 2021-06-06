package server

import (
	"context"
	"sync"

	"github.com/kjunn2000/chat-app/chat-api/internal/straper/data"
	proto "github.com/kjunn2000/chat-app/chat-api/proto"
	"go.uber.org/zap"
)

type Connection struct {
	stream proto.Chat_CreateStreamServer
	id     string
	active bool
	err    chan error
}

type ChatServer struct {
	Connections []*Connection
	Log         *zap.Logger
	ms          *data.MessageStore
	proto.UnimplementedChatServer
}

func NewChatServer(conn []*Connection, log *zap.Logger, ms *data.MessageStore) *ChatServer {

	s := &ChatServer{
		Connections: conn,
		Log:         log,
		ms:          ms,
	}
	return s
}

func (server *ChatServer) CreateStream(con *proto.Connect, stream proto.Chat_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     con.User.Id,
		active: con.Active,
		err:    make(chan error),
	}
	server.Connections = append(server.Connections, conn)
	server.Log.Info("Successful create a stream for client", zap.String("name", con.User.Name))

	return <-conn.err
}

func (server *ChatServer) BroadcastMessage(ctx context.Context, msm *proto.Message) (*proto.Close, error) {

	wg := sync.WaitGroup{}
	done := make(chan int)

	for _, v := range server.Connections {

		wg.Add(1)

		go func(msm *proto.Message, v *Connection) {
			defer wg.Done()
			if v.active && v.id != msm.Id {
				dberr := server.ms.CreateMessage(msm)
				if dberr != nil {
					server.Log.Fatal("Cannot save to db")
				}
				err := v.stream.Send(msm)
				if err == nil {
					server.Log.Info("Successful send a message to ", zap.String("id", v.id))

				} else {
					server.Log.Error("Error with stream , Failed when sending ", zap.String("id", v.id))
					v.active = false
					v.err <- err
				}
			}
		}(msm, v)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done

	return &proto.Close{}, nil
}
