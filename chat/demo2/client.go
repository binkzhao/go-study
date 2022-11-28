package main

import (
	"bufio"
	"fmt"
	"go-study/chat/demo2/core"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial(core.HOST_NETWORK, core.HOST_NAME)
	defer conn.Close()
	core.HandleErr(err)

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // 从服务端读取数据
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	ch := make(chan string)
	go core.WriteToConn(conn, ch)

	cin := bufio.NewScanner(os.Stdin)
	for cin.Scan() {
		switch cin.Text() {
		case core.OPERATE_REGISTER:
			ch <- core.OPERATE_REGISTER
			var userName, password, tel string
			fmt.Print("请输入用户名：")
			if cin.Scan() {
				userName = cin.Text()
			}
			fmt.Print("请输入密码：")
			if cin.Scan() {
				password = cin.Text()
			}
			fmt.Print("请输入手机号：")
			if cin.Scan() {
				tel = cin.Text()
			}
			ch <- userName
			ch <- password
			ch <- tel

		case core.OPERATE_LOGIN:
			ch <- core.OPERATE_LOGIN
			var userName, password string
			fmt.Print("请输入用户名：")
			if cin.Scan() {
				userName = cin.Text()
			}
			fmt.Print("请输入密码：")
			if cin.Scan() {
				password = cin.Text()
			}
			ch <- userName
			ch <- password

		case core.OPERATE_LOGOFF:
			ch <- core.OPERATE_LOGOFF

		case core.OPERATE_EXIT:
			os.Exit(0)

		case core.OPERATE_ADD:
			var userName string
			fmt.Print("请输入要添加好友的用户名：")
			if cin.Scan() {
				userName = cin.Text()
			}
			ch <- core.OPERATE_ADD
			ch <- userName

		case core.OPERATE_DELETE:
			fmt.Print("请输入要删除好友的用户名：")
			var userName string
			if cin.Scan() {
				userName = cin.Text()
			}
			ch <- core.OPERATE_DELETE
			ch <- userName

		case core.OPERATE_LIST:
			ch <- core.OPERATE_LIST

		case core.OPERATE_SENDTO:
			var userName, content string
			fmt.Print("请输入好友用户名：")
			if cin.Scan() {
				userName = cin.Text()
			}
			fmt.Print("请输入消息：")
			if cin.Scan() {
				content = cin.Text()
			}
			ch <- core.OPERATE_SENDTO
			ch <- userName
			ch <- content

		case core.OPERATE_SENDALL:
			var content string
			fmt.Print("请输入消息：")
			if cin.Scan() {
				content = cin.Text()
			}
			ch <- core.OPERATE_SENDALL
			ch <- content

		default:
			ch <- "none"
		}
	}

	<-done
}
