package core

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"io"
)

type UserService struct{}

func (userService UserService) Login(cin *bufio.Scanner) *UserModel {
	var userName, password string
	if cin.Scan() {
		userName = cin.Text()
	}
	if cin.Scan() {
		password = cin.Text()
	}

	file, err := os.Open(FILE_PATH_USER)
	defer file.Close()
	if err != nil {
		fmt.Println("func register: open ", FILE_PATH_USER)
	}

	for {
		var name, pwd, phone string
		_, err := fmt.Fscan(file, &name, &pwd, &phone) // 读取三个string
		if err != nil {
			fmt.Println(err)
			break
		}

		if userName == name && pwd == password {
			return &UserModel{UserName: userName}
		}
	}

	return nil
}

//读取文件user.txt，判断手机号和用户名是否重复注册
//将注册信息写入user.txt
//如果成功就创建messages、logs、friends等.txt文件
func (userService UserService) Register(cin *bufio.Scanner) bool {
	var userName, password, tel string
	if cin.Scan() {
		userName = cin.Text()
	}
	if cin.Scan() {
		password = cin.Text()
	}
	if cin.Scan() {
		tel = cin.Text()
	}

	file, err := os.OpenFile(FILE_PATH_USER, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println("Err in func register: ", FILE_PATH_USER)
	}

	var name, pwd, phone string
	hasExists := false
	for {
		_, err := fmt.Fscan(file, &name, &pwd, &phone) // 读取三个string
		if err != nil  && err != io.EOF {
			fmt.Println("Read User file fail: ", err)
			break
		}

		if err == io.EOF {
			break
		}

		if userName == name {
			hasExists = true
		}
	}

	if hasExists {
		return false
	}

	_, err = file.Write([]byte("\n" + userName + " " + password + " " + tel))
	HandleErr(err)
	userService.createInfo(userName)

	return true
}

func (userService UserService) createInfo(userName string) {
	file, err := os.OpenFile("friendsOf"+userName+".txt", os.O_CREATE, 0664)
	if err != nil {
		log.Fatalln("in func createInfo: ---")
	}
	file.Close()
	file, err = os.OpenFile("messagesOf"+userName+".txt", os.O_CREATE, 0664)
	file.Close()
	file, err = os.OpenFile("logsOf"+userName+".txt", os.O_CREATE, 0664)
	file.Close()
}
