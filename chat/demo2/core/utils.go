package core

import (
	"net"
	"log"
	"fmt"
)

// 写数据给服务端
func WriteToConn(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal("Err: ", err)
	}
}
