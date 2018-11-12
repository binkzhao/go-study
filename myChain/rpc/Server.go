package main

import (
	"encoding/json"
	"fmt"
	"go/myChain/core"
	"io"
	"net/http"
)

var blockChain *core.BlockChain

func run() {
	http.HandleFunc("/block-chain/get", blockChainGetHandler)
	http.HandleFunc("/block-chain/write", blockChainWriteHandler)
	fmt.Println("启动server, localhost:8881")
	http.ListenAndServe("localhost:8882", nil)
}

func blockChainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockChain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(bytes))
}

func blockChainWriteHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockChain.SendData(blockData)
	blockChainGetHandler(w, r)
}

func main() {
	blockChain = core.NewBlockChain()
	run()
}
