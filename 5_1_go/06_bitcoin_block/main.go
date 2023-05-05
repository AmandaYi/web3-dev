package main

func main() {
	blockChain := NewBlockChain("zzy")
	var cli CLI = CLI{BC: blockChain}
	cli.Run()
}
