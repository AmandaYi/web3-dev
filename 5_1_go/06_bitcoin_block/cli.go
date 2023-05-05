package main

import (
	"fmt"
	"os"
)

type CLI struct {
	BC *BlockChain
}

const Usage = `
	addBlock --data "add data to blockchain"
	printChain 			 "print all blockchain data"
	getBalance --address "address"
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
				//data := args[3]
				//b. 使用bc添加区块AddBlock
				//cli.AddBlock(data)
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
	case "getBalance":
		{
			if len(args) == 4 && args[2] == "--address" {
				data := args[3]
				cli.GetBalance(data)
			} else {
				fmt.Printf("余额参数使用不当，请检查")
				fmt.Printf(Usage)
			}
		}

	default:
		fmt.Println("无效命令")
	}
}
