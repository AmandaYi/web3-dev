package main

func main() {
	blockChain := NewBlockChain()

	blockChain.AddBlock("小明向小红转了一笔钱")
	blockChain.PrintAll()

	blockChain.AddBlock("小李向小张转了一笔钱")
	blockChain.PrintAll()
}
