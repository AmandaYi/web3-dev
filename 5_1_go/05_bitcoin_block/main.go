package main

import "05_bitcoin_block/blockCore"

func main() {
	blockChain := blockCore.NewBlockChain()
	var cli CLI = CLI{BC: blockChain}
	cli.Run()
}
