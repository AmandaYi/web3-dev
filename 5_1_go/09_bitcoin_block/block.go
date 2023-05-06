package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
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
	//Data []byte
	//区块交易体
	Transactions []*Transaction

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
func GenesisBlock(address string, data string, preHash []byte) *Block {
	var txs []*Transaction = []*Transaction{NewCoinBaseTX(address, data)}
	return NewBlock(txs, preHash)
}
func NewBlock(txs []*Transaction, preHash []byte) *Block {
	//这里不需要验证交易，因为这是创世块，也是矿机池，在交易函数里面是直接通过的
	var block = Block{
		PreHash: preHash,
		Hash:    []byte{},

		//Data:    data,
		Transactions: txs, // 交易体

		//--- 新添加的
		Version:    00,
		MerKleRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 100,
		Nonce:      100,
	}
	// 设置自己的Hash值
	//blockCore.SetHash()

	//修改为通过系统计算得到比系统预设hash小的hash值，刚好可以得到nonce随机值
	block.Hash, block.Nonce = NewProofOfWork(&block).Run()
	return &block
}

func (b *Block) PrintAll() {
	fmt.Println("---STR---")
	fmt.Printf("blockCore.PreHash: %x\n", b.PreHash)
	//fmt.Printf("blockCore.Data: %s\n", string(b.Data))
	fmt.Println("---END---")
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
		b.PreHash,
		//b.Data,
		uint64ToByte(b.Version),
		b.MerKleRoot,
		uint64ToByte(b.TimeStamp),
		uint64ToByte(b.Difficulty),
		uint64ToByte(b.Nonce),
	}
	tmp := bytes.Join(genesisInfoByte, []byte(""))
	// 对临时变量进行加密
	result := sha256.Sum256(tmp)

	// 把结果作为自己的Hash存起来
	b.Hash = result[:]
}

//设置交易的梅克尔根,其实是个平衡二叉树组成的hash,这里简单写,直接写hash
func (b *Block) SetTransactionsHash() []byte {
	var result [32]byte

	var tmpResult [][]byte

	for _, v := range b.Transactions {
		tmpResult = append(tmpResult, v.TXID)
	}

	data := bytes.Join(tmpResult, []byte{})

	result = sha256.Sum256(data)

	return result[:]
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

//区块序列化函数
// 使用序列化的函数，会把结构体的[]byte还原成nil，造成隐性错误,跟下面的Deserialize无关
func (b *Block) Serialize() []byte {
	//将block数据转换成字节流
	var buffer bytes.Buffer
	//创建一个序列化编码器
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(b)
	if err != nil {
		fmt.Println("编码失败", err)
		panic(err)
	}
	return buffer.Bytes()
}

//区块反序列化函数

func (b *Block) Deserialize(data []byte) Block {
	var block Block
	var buffer bytes.Buffer
	//将data写入buffer
	_, err := buffer.Write(data)
	if err != nil {
		fmt.Println("buffer写入失败", err)
		panic(err)
	}

	//创建一个反序列化解码器，用于解码
	decoder := gob.NewDecoder(&buffer)
	//将buffer数据转换成block
	err = decoder.Decode(&block)
	if err != nil {
		fmt.Println("解码失败", err)
		panic(err)
	}
	return block
}
