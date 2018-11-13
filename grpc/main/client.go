package main

import (
	"google.golang.org/grpc"
	"log"
	pb "go/grpc/protos"
	"os"
	"golang.org/x/net/context"
	"go/grpc/consts"
	"fmt"
	"strings"
)

func main() {
	conn, err := grpc.Dial(consts.HOST, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)
	res, err := client.Echo(context.Background(), &pb.HelloReq{
		Msg: strings.Join(os.Args[1:], " "),
	})

	if err != nil {
		fmt.Errorf("client echo failed. err: [%v]", err)
		return
	}

	fmt.Printf("Msg from server: %s", res.GetMsg())
}