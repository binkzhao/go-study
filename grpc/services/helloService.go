package services

import (
	"context"
	"go/grpc/protos"
	"fmt"
)

type HelloService struct {}

func (this *HelloService) Echo(ctx context.Context, req *hello.HelloReq) (*hello.HelloRes, error) {
	fmt.Printf("message from client: %s\n", req.GetMsg())
	res := &hello.HelloRes{
		Msg: "Server already receive your msg:" + req.GetMsg(),
	}

	return res, nil
}
