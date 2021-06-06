package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kjunn2000/chat-app/chat-api/interceptors"
	proto "github.com/kjunn2000/chat-app/chat-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var log grpclog.LoggerV2
var client proto.ChatClient
var wg *sync.WaitGroup
var scanner *bufio.Scanner

func init() {
	log = grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func createStream(user *proto.User) error {

	var streamErr error

	connect := &proto.Connect{
		User:   user,
		Active: true,
	}

	stream, err := client.CreateStream(context.Background(), connect)
	if err != nil {
		log.Fatalf("Cannot establish stream at client side ==> %v", err)
	}

	wg.Add(1)
	go func(stream proto.Chat_CreateStreamClient) {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Infof("GG cannot receive message %v", err)
				streamErr = fmt.Errorf("GG cannot receive message %v", err)
				break
			}
			fmt.Printf("\n%v : %s\n", res.Id, res.Content)
		}

	}(stream)
	return streamErr
}

func sendMessage(user *proto.User) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			scanner.Scan()
			msg := &proto.Message{
				Id:        user.Id,
				Content:   scanner.Text(),
				Timestamp: time.Now().String(),
			}
			_, err := client.BroadcastMessage(context.Background(), msg)
			if err != nil {
				log.Fatalf("Unable to sent mesasge : %v", err)
			}
		}
	}()

}

func main() {
	done := make(chan int)
	wg = &sync.WaitGroup{}

	scanner = bufio.NewScanner(os.Stdout)
	fmt.Print("Username ", " : ")
	scanner.Scan()

	nf := flag.String("n", scanner.Text(), "Flag for name")
	flag.Parse()

	// id := sha256.Sum256([]byte(time.Now().String() + *nf))

	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure(),
		grpc.WithStreamInterceptor(interceptors.AddToStreamRequest),
		grpc.WithUnaryInterceptor(interceptors.AddToUnaryReqeust),
	)

	if err != nil {
		log.Fatalf("the connection is established unsuccessful => %v", err)
	}
	client = proto.NewChatClient(conn)

	User := proto.User{
		// Id:   hex.EncodeToString(id[:]),
		Id:   uuid.NewString(),
		Name: *nf,
	}

	createStream(&User)

	sendMessage(&User)

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done
}
