package main

import "fmt"

// 区块链 - 链表
type BlockChain struct {
	Blocks []*Block
}

// 遍历区块链
func (b *BlockChain) PrintAll() {
	fmt.Println("--------BlockChain--STR--------")
	for _, v := range b.Blocks {
		v.PrintAll()
	}
	fmt.Println("--------BlockChain--END--------")
}

// 初始化区块链 - 链表
func NewBlockChain() *BlockChain {
	var bc BlockChain
	// 初始化第一个区块
	const genesisInfo string = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
	firstBlock := GenesisBlock(genesisInfo, []byte{})
	//把第一个区块放入链表即可，形成一个新的链表
	bc.Blocks = []*Block{firstBlock}
	return &bc
}

// 追加区块链数据
func (b *BlockChain) AddBlock(data string) {
	// 初始化块信息
	var newBlock Block
	// 得到上一个区块的Hash的值，作为自己的preHash值
	var lenBlocks int = len(b.Blocks)
	lastBlock := b.Blocks[lenBlocks-1]
	newBlock.PreHash = lastBlock.Hash
	// 保存data信息
	newBlock.Data = []byte(data)
	// 计算自己的Hash值
	newBlock.SetHash()
	// 追加到区块链里面
	b.Blocks = append(b.Blocks, &newBlock)
}
