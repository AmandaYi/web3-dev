package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

// 创建一个基本版本的区块
type Block struct {
	// 叔父Hash
	PreHash []byte
	//当前hash
	Hash []byte
	//区块数据
	Data []byte

	//---- 新添加的值
	// 版本号
	Version uint64
	// 梅克尔根,本身是个Hash值
	MerKleRoot []byte

	// 时间戳
	TimeStamp uint64
	// 难度值(根据挖矿时间得出)
	Difficulty uint64
	// 随机数 nonce
	Nonce uint64
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
		//--- 新添加的
		Version:    00,
		MerKleRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 100,
		Nonce:      100,
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
	/*
		// 把当前区块的全部数据，存到一个临时变量
		var genesisInfoByte []byte
		genesisInfoByte = append(genesisInfoByte, b.PreHash...)
		genesisInfoByte = append(genesisInfoByte, b.Data...)
		//新添加的
		genesisInfoByte = append(genesisInfoByte, uint64ToByte(b.Version)...)
		genesisInfoByte = append(genesisInfoByte, b.MerKleRoot...)
		genesisInfoByte = append(genesisInfoByte, uint64ToByte(b.TimeStamp)...)
		genesisInfoByte = append(genesisInfoByte, uint64ToByte(b.Difficulty)...)
		genesisInfoByte = append(genesisInfoByte, uint64ToByte(b.Nonce)...)
	*/
	genesisInfoByte := [][]byte{
		b.PreHash, b.Data, uint64ToByte(b.Version), b.MerKleRoot, uint64ToByte(b.TimeStamp), uint64ToByte(b.Difficulty), uint64ToByte(b.Nonce),
	}
	tmp := bytes.Join(genesisInfoByte, []byte(""))
	// 对临时变量进行加密
	result := sha256.Sum256(tmp)

	// 把结果作为自己的Hash存起来
	b.Hash = result[:]
}

// 辅助函数
func uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}
