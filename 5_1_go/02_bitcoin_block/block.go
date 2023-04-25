package main

import (
	"crypto/sha256"
	"fmt"
)

// 创建一个基本版本的区块
type Block struct {
	// 叔父Hash
	PreHash []byte
	//当前hash
	Hash []byte
	//区块数据
	Data []byte
}

// 使用函数初始化结构体，产生创世块
func GenesisBlock(data string, preHash []byte) *Block {
	return NewBlock([]byte(data), preHash)
}
func NewBlock(data []byte, preHash []byte) *Block {
	var block = Block{
		PreHash: preHash,
		Hash:    []byte{},
		Data:    data,
	}
	// 设置自己的Hash值
	block.SetHash()
	return &block
}

func (b *Block) PrintAll() {
	fmt.Printf("preHash:%x\n", b.PreHash)
	fmt.Printf("Hash:%x\n", b.Hash)
	fmt.Println("Data:", string(b.Data))
}

// SetHash
func (b *Block) SetHash() {
	// 把当前区块的全部数据，存到一个临时变量
	var genesisInfoByte []byte
	genesisInfoByte = append(genesisInfoByte, b.PreHash...)
	genesisInfoByte = append(genesisInfoByte, b.Data...)
	// 对临时变量进行加密
	result := sha256.Sum256(genesisInfoByte)

	// 把结果作为自己的Hash存起来
	b.Hash = result[:]
}
