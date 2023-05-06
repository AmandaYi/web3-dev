package main

func main() {
	blockChain := NewBlockChain("LwxuwY91mK5hnAtQHx3neKagrcXqEHvDuZ")
	var cli CLI = CLI{BC: blockChain}
	cli.Run()
}
