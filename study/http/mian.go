package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello World!")
	})
	http.ListenAndServe("127.0.0.1:9900", nil)
}
