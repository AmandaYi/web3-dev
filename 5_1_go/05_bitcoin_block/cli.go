package main

import (
	"05_bitcoin_block/blockCore"
	"fmt"
	"os"
)

type CLI struct {
	BC *blockCore.BlockChain
}

const Usage = `
	addBlock --data DATA "add data to blockchain"
	printChain 			 "print all blockchain data"
`

func (cli *CLI) Run() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}
	//处理命令
	cmd := args[1]
	switch cmd {
	case "addBlock":
		{
			//确保命令有效
			if len(args) == 4 && args[2] == "--data" {
				//获取命令的数据
				//a. 获取数据
				data := args[3]
				//b. 使用bc添加区块AddBlock
				cli.AddBlock(data)
			} else {
				fmt.Printf("添加区块参数使用不当，请检查")
				fmt.Printf(Usage)
			}
		}
	case "printChain":
		{
			fmt.Println("print")
			cli.PrintBlockChain()
		}
	default:
		fmt.Println("无效命令")
	}
}
