package main

import (
	"net/rpc"
	"net"
	"log"
	"go/study/rpc/core"
)

func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rect := new(core.Rect)

	rpc.Register(rect)

	//获取tcpAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", core.Remote)
	chkError(err)

	// 监听端口
	tcpListen, err := net.ListenTCP("tcp", tcpAddr)
	chkError(err)
	// 死循环处理连接请求
	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			continue
		}

		// 使用goroutine单独处理rpc连接请求
		go rpc.ServeConn(conn)
	}
}