package main

func main() {
	blockChain := NewBlockChain("zzy")
	var cli CLI = CLI{BC: blockChain}
	cli.Run()
	//cli.Send("zzy", "zy", 2, "大佬的挖矿机器", "zzy转zy2")
	//cli.GetBalance("zzy")
	//cli.CreateWallet()
}
