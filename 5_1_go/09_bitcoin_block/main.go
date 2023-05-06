package main

func main() {
	blockChain := NewBlockChain("LwxuwY91mK5hnAtQHx3neKagrcXqEHvDuZ")
	var cli CLI = CLI{BC: blockChain}
	cli.Run()
	//cli.Send("LwxuwY91mK5hnAtQHx3neKagrcXqEHvDuZ", "M4Pam9Gemu3WJmNh1j7fTf8iH261iYvxeN", 2, "LvqKBZKNNqYs5G8c5nTPcRsf97TiqRNXre", "2")
}
