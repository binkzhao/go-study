package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  // 区块编号
	Timestamp     int64  // 区块时间戳
	PrevBlockHash string // 上一个区块的哈希值
	Hash          string // 当前区块的哈希值
	Data          string // 数据
}

// 计算区块的哈希值
func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInByte := sha256.Sum256([]byte(blockData))
	hashStr := hex.EncodeToString(hashInByte[:])
	return hashStr
}

// 生成一个新的区块
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

// 创建创世区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}
