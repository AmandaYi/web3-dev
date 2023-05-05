package main

import (
	"05_bitcoin_block/blockCore"
	"fmt"
)

func (c *CLI) AddBlock(data string) {
	c.BC.AddBlock(data)
	fmt.Println("添加区块成功")
}
func (c *CLI) PrintBlockChain() {
	blockChain := c.BC
	it := blockCore.NewBlockChainIterator(blockChain)
	for {
		blockCore := it.Next()
		fmt.Printf("===========================\n\n")
		fmt.Printf("版本号: %d\n", blockCore.Version)
		fmt.Printf("前区块哈希值: %x\n", blockCore.PreHash)
		fmt.Printf("梅克尔根: %x\n", blockCore.MerKleRoot)
		fmt.Printf("时间戳: %d\n", blockCore.TimeStamp)
		fmt.Printf("难度值(随便写的）: %d\n", blockCore.Difficulty)
		fmt.Printf("随机数 : %d\n", blockCore.Nonce)
		fmt.Printf("当前区块哈希值: %x\n", blockCore.Hash)
		fmt.Printf("区块数据 :%s\n", blockCore.Data)

		if len(blockCore.PreHash) == 0 {
			fmt.Printf("区块链遍历结束！")
			break
		}
	}
}
