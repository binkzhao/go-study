package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"bytes"
	"log"
)

func main() {
	log.Fatal()
}

func demo1()  {
	// 1. GET
	res, _ := http.Get("http://baidu.com")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	// 2. POST
	postBody := "{\"action\":20}"
	res, err := http.Post("http://baidu.com", "application/json;charset=utf-8", bytes.NewBuffer([]byte(postBody)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}
