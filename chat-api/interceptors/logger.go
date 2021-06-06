package interceptors

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func LogUnaryRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (response interface{}, err error) {
	fmt.Printf("[Unary] Request for : %s\n", info.FullMethod)
	return handler(ctx, req)
}

func LogStreamRequst(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	fmt.Printf("[Stream] Request for : %s\n", info.FullMethod)
	handler(srv, ss)
	return nil
}

func AddToUnaryReqeust(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	fmt.Printf("[Unary] Request from: %s\n", method)
	invoker(ctx, method, req, reply, cc, opts...)
	return nil
}

func AddToStreamRequest(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Printf("[Stream] Request from: %s\n", method)
	clientStream, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		fmt.Printf("Sysem error : %s", err.Error())
	}
	return clientStream, err
}
