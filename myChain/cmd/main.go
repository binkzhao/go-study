package main

import "go-study/myChain/core"

func main() {
	blockChain := core.NewBlockChain()
	blockChain.SendData("send 1 BTC to Amy")
	blockChain.SendData("send 100 BTC to Bink")
	blockChain.SendData("send 1 BTC to Jack")
	blockChain.Print()
}
