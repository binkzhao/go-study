package main

import (
	pb "go/grpc/protos"
	"net"
	"log"
	"google.golang.org/grpc"
	"go/grpc/consts"
	"fmt"
	"go/grpc/services"
)

func main() {
	listener, err := net.Listen("tcp", consts.PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &services.HelloService{}) // 注册hello service

	fmt.Println("Starting server on port ", consts.PORT)
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
