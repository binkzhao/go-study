package main

import (
	"bufio"
	"fmt"
	"go-study/chat/demo2/core"
	"log"
	"net"
)

type userCH chan<- string

var (
	messages = make(chan string)
	entering = make(chan userCH)
	leaving  = make(chan userCH)
	clients  = make(map[string]userCH)
)

func main() {
	listener, err := net.Listen(core.HOST_NETWORK, core.HOST_NAME)
	core.HandleErr(err)
	fmt.Println("server is running...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	var client *core.UserModel
	ch := make(chan string)
	go core.WriteToConn(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "Welcome! " + who
	ch <- core.USAGE
	ch <- ">>"
	input := bufio.NewScanner(conn)
	userService := core.UserService{}

	for input.Scan() {
		var op = input.Text()
		fmt.Println("Operate： ", op)
		switch op {
		case core.OPERATE_REGISTER:
			ok := userService.Register(input)
			if !ok {
				ch <- "Fail register=_=||\nmaybe your userName or phoneNumber is invalid"
			} else {
				ch <- "Success register!"
			}

		case core.OPERATE_LOGIN:
			client = userService.Login(input)
			if client != nil {
				clients[client.UserName] = ch
				ch <- "Success login!\n" + "Your new messages:"
				client.ToRead(ch)
			} else {
				ch <- "Fail login=_=||maybe your userName or password is wrong,please check them carefully:)"
			}
		}
	}
}
