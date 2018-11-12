package core

import (
	"fmt"
	"log"
)

// 区块链结构体
type BlockChain struct {
	Blocks []*Block
}

// 创建区块链
func NewBlockChain() *BlockChain {
	genesisBlock := GenerateGenesisBlock()
	blockChain := BlockChain{}
	blockChain.AppendBlock(&genesisBlock)
	return &blockChain
}

func (self *BlockChain) SendData(data string) {
	preBlock := self.Blocks[len(self.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	self.AppendBlock(&newBlock)
}

// 添加一个新的区块
func (self *BlockChain) AppendBlock(newBlock *Block) {
	if len(self.Blocks) == 0 {
		self.Blocks = append(self.Blocks, newBlock)
		return
	}

	if self.isValidBlock(*newBlock) {
		self.Blocks = append(self.Blocks, newBlock)
	} else {
		log.Fatalln("invalid Block")
	}
}

// 验证一个新的区块是否有效
func (self *BlockChain) isValidBlock(newBlock Block) bool {
	oldBlock := self.Blocks[len(self.Blocks)-1]

	if newBlock.Index-1 != oldBlock.Index {
		return false
	}

	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// 显示区块信息
func (self *BlockChain) Print() {
	for _, block := range self.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Prev.Hash: %s\n", block.PrevBlockHash)
		fmt.Printf("Curr.Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Println()
	}
}
