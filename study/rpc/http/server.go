package main

import (
	"fmt"
	"go-study/study/rpc/core"
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	rect := new(core.Rect)
	rpc.Register(rect) // 注册一个rect服务
	rpc.HandleHTTP()   // 把服务处理绑定到http协议上
	fmt.Println("run server on port 8000...")
	err := http.ListenAndServe(core.Remote, nil)
	if err != nil {
		log.Fatal(err)
	}
}
