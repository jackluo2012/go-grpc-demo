package main

import (
	"net"

	"golang.org/x/net/context"

	pb "hello/proto/hello_http"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloHTTPService并实现约定的接口
type helloHTTPService struct{}

// HelloHTTPService 实现服务端接口
var HelloHTTPService = helloHTTPService{}

// SayHello ...
func (h helloHTTPService) SayHello(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	resp := new(pb.HelloHTTPResponse)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloHTTPService
	pb.RegisterHelloHTTPServer(s, HelloHTTPService)

	fmt.Println("Listen on " + Address)

	s.Serve(listen)
}
