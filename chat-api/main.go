package main

import (
	"net"

	"github.com/kjunn2000/chat-app/chat-api/interceptors"
	"github.com/kjunn2000/chat-app/chat-api/internal/straper/data"
	cs "github.com/kjunn2000/chat-app/chat-api/internal/straper/server"
	proto "github.com/kjunn2000/chat-app/chat-api/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {

	log, err := zap.NewDevelopment()
	defer log.Sync()
	if err != nil {
		log.Fatal("Failed to create logger")
	}
	ms := data.NewMessageStore(log)
	s := cs.NewChatServer(make([]*cs.Connection, 0), log, ms)
	uic := grpc.ChainUnaryInterceptor(interceptors.LogUnaryRequest)
	sic := grpc.ChainStreamInterceptor(interceptors.LogStreamRequst)
	grpcServer := grpc.NewServer(uic, sic)
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("Unable to start server at port 9090")
	}
	log.Info("Staring server : ", zap.Int("port", 9090))
	proto.RegisterChatServer(grpcServer, s)
	grpcServer.Serve(listen)
}
