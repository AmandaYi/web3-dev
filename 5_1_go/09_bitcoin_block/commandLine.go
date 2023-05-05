package main

import (
	"fmt"
)

func (c *CLI) PrintBlockChain() {
	blockChain := c.BC
	it := NewBlockChainIterator(blockChain)
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
		//fmt.Printf("区块数据 :%s\n", blockCore.Data)
		fmt.Printf("区块数据 :%s\n", blockCore.Transactions[0].TXInputs[0].ScriptSig)
		if len(blockCore.PreHash) == 0 {
			fmt.Printf("区块链遍历结束！")
			break
		}
	}
}

func (c *CLI) GetBalance(address string) {
	utxos := c.BC.FindUTXOs(address)
	var total float64
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("%s的余额是%f", address, total)
}

func (c *CLI) Send(from string, to string, amount float64, miner string, data string) {
	fmt.Printf("from: %s, to: %s, amount: %f, miner: %s, data: %s\n", from, to, amount, miner, data)

	//创建普通交易
	normalTransaction := NewTransaction(from, to, amount, c.BC)

	//创建挖矿交易
	coinTransaction := NewCoinBaseTX(miner, data)
	//创建一个区块
	//把交易存到区块里面，同时把区块保存到区块链
	c.BC.AddBlock([]*Transaction{coinTransaction, normalTransaction})

}
func (c *CLI) CreateWallet() {
	//使用钱包容器创建
	ww := NewWalletWrapper()
	address := ww.CreateWallet()

	//直接操作指定的钱包测试
	//wallet := NewWallet()
	//address := wallet.GenAddress()
	//fmt.Println("钱包私钥", wallet.PrivateKey)
	//fmt.Println("钱包公钥", wallet.PublicKey)

	fmt.Println("地址", address)
}

func (c *CLI) ListAddress() {
	ww := NewWalletWrapper()
	ww.GetListWallet()
}
