package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Person struct {
	Name string
	Age  int
}

var ws = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	indexFile, err := os.Open("html/index.html")
	checkErr(err)

	index, err := ioutil.ReadAll(indexFile)
	checkErr(err)

	http.HandleFunc("/websocket", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := ws.Upgrade(writer, request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}

			if string(msg) == "ping" {
				fmt.Println("ping")
				time.Sleep(time.Second * 2)
				err = conn.WriteMessage(msgType, []byte("pong"))
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				conn.Close()
				fmt.Println(string(msg))
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(index))
	})

	fmt.Println("running service.")
	err = http.ListenAndServe(":3000", nil)
	fmt.Println(err)
}

func checkErr(err error) {
	if err != nil {
		panic(any(err))
	}
}
