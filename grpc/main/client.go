package main

import (
	"google.golang.org/grpc"
	"log"
	pb "go/grpc/protos"
	"os"
	"golang.org/x/net/context"
	"math/rand"
	"go/grpc/consts"
	"encoding/json"
	"fmt"
)

const defaultName = "Bink"

func main() {
	conn, err := grpc.Dial(consts.HOST, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewStudentServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.StudentAdd(
		context.Background(),
		&pb.Request{
			Sid: int32(rand.Int()%10000000),
			Name: name,
			Age: int32(rand.Int()%100),
		},
	)
	if err != nil {
		log.Fatalf("Could not Add Student: %v", err)
	}

	repStr,_ := json.Marshal(r)
	fmt.Println("Response: ", string(repStr))
}