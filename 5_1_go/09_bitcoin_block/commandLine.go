package main

import (
	"fmt"
)

func (c *CLI) PrintBlockChain() {
	blockChain := c.BC
	it := NewBlockChainIterator(blockChain)
	for {
		blockCore := it.Next()
		for _, tx := range blockCore.Transactions {
			fmt.Println(tx)
		}
		//fmt.Printf("===========================\n\n")
		//fmt.Printf("版本号: %d\n", blockCore.Version)
		//fmt.Printf("前区块哈希值: %x\n", blockCore.PreHash)
		//fmt.Printf("梅克尔根: %x\n", blockCore.MerKleRoot)
		//fmt.Printf("时间戳: %d\n", blockCore.TimeStamp)
		//fmt.Printf("难度值(随便写的）: %d\n", blockCore.Difficulty)
		//fmt.Printf("随机数 : %d\n", blockCore.Nonce)
		//fmt.Printf("当前区块哈希值: %x\n", blockCore.Hash)
		//fmt.Printf("区块数据 :%s\n", blockCore.Data)
		//fmt.Printf("区块数据 :%s\n", blockCore.Transactions[0].TXInputs[0].ScriptSig)
		if len(blockCore.PreHash) == 0 {
			fmt.Printf("区块链遍历结束！")
			break
		}
	}
}

func (c *CLI) GetBalance(address string) {
	if !IsValidAddress(address) {
		fmt.Printf("不是合法有效的地址: %s \n", address)
		return
	}
	//获取余额
	//获取余额需要指定地址，通过遍历整个账本，从而找到这个地址可用的utxo，为此我们要做两件事：
	//1.  校验地址的有效性
	//传递过来的地址有可能是无效的，无效的地址直接返回即可。
	//2.  逆推出公钥哈希
	//并不是所有的地址都是本地生成的，有可能是别人的地址，所以我们需要逆推而不是打开钱包去
	// 修改交易结构->遍历账本->调用FindUTXOs函数

	//这里传入的是随机一个地址，去获取余额，因此，需要逆推到公钥，而不是找钱包拿公钥，因为这个地址很可能就不是自己的
	publicKeyHash := GetPublicKeyHashFromAddress(address)

	//utxos := c.BC.FindUTXOs(address)
	utxos := c.BC.FindUTXOs(publicKeyHash)
	var total float64
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("%s的余额是%f", address, total)
}

func (c *CLI) Send(from string, to string, amount float64, miner string, data string) {
	if !IsValidAddress(from) {
		fmt.Printf("不是合法有效的地址: %s \n", from)
		return
	}
	if !IsValidAddress(to) {
		fmt.Printf("不是合法有效的地址: %s \n", to)
		return
	}
	if !IsValidAddress(miner) {
		fmt.Printf("不是合法有效的地址: %s \n", miner)
		return
	}
	fmt.Printf("from: %s, to: %s, amount: %f, miner: %s, data: %s\n", from, to, amount, miner, data)

	//创建普通交易
	normalTransaction := NewTransaction(from, to, amount, c.BC)
	if normalTransaction == nil {
		return
	}

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
