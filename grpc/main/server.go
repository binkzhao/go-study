package main

import (
	"golang.org/x/net/context"
	pb "go/grpc/protos"
	"os"
	"net"
	"log"
	"google.golang.org/grpc"
	"go/grpc/consts"
	"fmt"
)

type Server struct{}

func (s *Server) StudentAdd(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	file, _ := os.OpenFile("stulist.txt", os.O_RDWR|os.O_CREATE, 0665)
	defer file.Close()
	_, err :=file.WriteString(string(req.Sid) + "\t " + req.Name + "\t " + string(req.Age))
	if err != nil {
		log.Fatalf("Add Student Err: %v", err)
		return &pb.Response{
			Code: 500,
			Msg:  "Add Student " + req.Name + " Failed.",
		}, nil
	}

	return &pb.Response{
		Code: 0,
		Msg:  "Add Student " + req.Name + " Success.",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", consts.PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStudentServiceServer(s, &Server{})

	fmt.Println("Starting server on port ", consts.PORT)
	s.Serve(listener)
}
