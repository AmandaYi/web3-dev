package main

import (
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	BC *BlockChain
}

const Usage = `
	printChain 			 "print all blockchain data"
	getBalance --address "address"
	send FROM TO AMOUNT MINER "DATA"
	createWallet "创建一个钱包"
	listAddress "listAddress"
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
	//addBlock --data "add data to blockchain
	//case "addBlock":
	//	{
	//		//确保命令有效
	//		if len(args) == 4 && args[2] == "--data" {
	//			//获取命令的数据
	//			//a. 获取数据
	//			//data := args[3]
	//			//b. 使用bc添加区块AddBlock
	//			//cli.AddBlock(data)
	//		} else {
	//			fmt.Printf("添加区块参数使用不当，请检查")
	//			fmt.Printf(Usage)
	//		}
	//	}
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
	case "send":
		{
			if len(args) == 7 {
				//send FROM TO AMOUNT MINER "DATA"
				from := args[2]
				to := args[3]
				amount, _ := strconv.ParseFloat(args[4], 64)
				miner := args[5]
				data := args[6]
				cli.Send(from, to, amount, miner, data)
				fmt.Println("Send")
			} else {
				fmt.Printf("SEND参数使用不当，请检查")
				fmt.Printf(Usage)
			}
		}
	case "createWallet":
		{
			cli.CreateWallet()
		}
	case "listAddress":
		{
			cli.ListAddress()
		}
	default:
		fmt.Println("无效命令")
	}
}
