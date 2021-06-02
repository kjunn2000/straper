package main

import (
	"context"
	"net"
	"os"
	"sync"

	proto "github.com/kjunn2000/chat-app/chat-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Connection struct {
	stream proto.Chat_CreateStreamServer
	id     string
	active bool
	err    chan error
}

type Server struct {
	Connections []*Connection
	log         grpclog.LoggerV2
	proto.UnimplementedChatServer
}

func (server *Server) CreateStream(con *proto.Connect, stream proto.Chat_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     con.User.Id,
		active: con.Active,
		err:    make(chan error),
	}
	server.Connections = append(server.Connections, conn)
	server.log.Infof("Successful create a stream for client %v", con.User.Name)

	return <-conn.err
}

func (server *Server) BroadcastMessage(ctx context.Context, msm *proto.Message) (*proto.Close, error) {

	wg := sync.WaitGroup{}
	done := make(chan int)

	for _, v := range server.Connections {

		wg.Add(1)

		go func(msm *proto.Message, v *Connection) {
			defer wg.Done()
			if v.active && v.id != msm.Id {
				err := v.stream.Send(msm)
				if err == nil {
					server.log.Info("Successful send a message to ", v.id)
				} else {
					server.log.Errorf("Error with stream - %v , Failed when sending to %v", v.stream, v.id)
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

func main() {

	log := grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	s := &Server{
		Connections: []*Connection{},
		log:         log,
	}
	grpcServer := grpc.NewServer()
	listen, err := net.Listen("tcp", ":9050")
	if err != nil {
		s.log.Fatalf("Unable to start server at port %v : ", "9050")
	}
	s.log.Infof("Starting server %v", "9050")
	proto.RegisterChatServer(grpcServer, s)
	grpcServer.Serve(listen)
}
