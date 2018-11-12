package core

import (
	"os"
	"log"
	"fmt"
)

// client model
type UserModel struct {
	UserName string
}

func (user *UserModel) ToRead(ch chan string) {
	msgFilePath := "messagesOf" + user.UserName + ".txt"
	logsFilePath := "logsOf" + user.UserName + ".txt"
	msgsFile, err := os.OpenFile(msgFilePath, os.O_RDWR, 0666) // RW
	defer msgsFile.Close()
	logsFile, err := os.OpenFile(logsFilePath, os.O_APPEND|os.O_WRONLY, 0666)
	defer logsFile.Close()
	if err != nil {
		log.Fatalln("func toRead, fail to open file")
	}

	for {
		// 格式 ： 时间、好友名字、信息内容
		var time, friendName, msg string
		_, err := fmt.Fscan(msgsFile, &time, &friendName, &msg)
		if err != nil {
			fmt.Println(err)
			break
		}

		temp := time + " " + friendName + " " + msg
		_, err = logsFile.Write([]byte(time + " " + friendName + " " + user.UserName + " " + msg + "\n"))
		if err != nil {
			fmt.Println("write logs wrong! in func toRead")
		}
		ch <-temp
	}

	msgsFile.Seek(0, 0)
	msgsFile.Truncate(0) // 清空未读信息
}
