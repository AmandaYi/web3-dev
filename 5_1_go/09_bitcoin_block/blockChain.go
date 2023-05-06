package main

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
)

const BLOCK_CHAIN_DB = "blockChain.db"
const BLOCK_BUCKET = "blockBucket"
const LASH_HASH = "lastHash"

// 区块链 - 链表
type BlockChain struct {
	//Blocks []*Block

	db *bolt.DB

	lastHash []byte
}

// 遍历区块链
func (b *BlockChain) PrintAll() {
	b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_BUCKET))
		//通过唯一key，找到最后一个hash值
		lastHash := bucket.Get([]byte(LASH_HASH))
		//fmt.Printf("lastHash %x\n", lastHash)
		var block Block
		var hash []byte = lastHash
		for {
			tmpBlock := bucket.Get(hash)
			if tmpBlock == nil {
				fmt.Println("没有更多数据了")
				break
			}
			block = block.Deserialize(tmpBlock)
			block.PrintAll()
			hash = block.PreHash
		}
		return nil
	})
}

// 初始化区块链 - 链表
func NewBlockChain(address string) *BlockChain {
	if !IsValidAddress(address) {
		fmt.Printf("不是合法有效的地址: %s \n", address)
		return nil
	}
	//var bc BlockChain
	//// 初始化第一个区块
	//const genesisInfo string = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
	//firstBlock := GenesisBlock(genesisInfo, []byte{})
	////把第一个区块放入链表即可，形成一个新的链表
	//bc.Blocks = []*Block{firstBlock}

	var db *bolt.DB
	var lastHash []byte
	var bc BlockChain
	var err error
	db, err = bolt.Open(BLOCK_CHAIN_DB, 0600, nil)
	if err != nil {
		panic("创建数据库失败")
	}

	//blockBucket
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_BUCKET))
		if bucket == nil {
			//代表没有，进行创建
			bucket, err = tx.CreateBucket([]byte(BLOCK_BUCKET))
			if err != nil {
				panic("创建bucket失败")
			}

			const genesisInfo string = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
			// 初始化一个基本块
			firstBlock := GenesisBlock(address, genesisInfo, []byte{})

			// 初始化一个基本交易体
			//创建coinbase交易 <---这⾥修改
			coinBaseTX := NewCoinBaseTX(address, genesisInfo)
			var txs []*Transaction = []*Transaction{coinBaseTX}
			// 设置块的交易体
			firstBlock.Transactions = txs

			//保存一份数据，
			bucket.Put(firstBlock.Hash, firstBlock.Serialize())

			//同时保持最后的hash值
			fmt.Printf("第一次新创建的 firstBlock.Hash %x\n", firstBlock.Hash)
			bucket.Put([]byte(LASH_HASH), firstBlock.Hash)
			lastHash = firstBlock.Hash
		} else {
			//如果已经有bucket了，那就读取lastHash即可
			lastHash = bucket.Get([]byte(LASH_HASH))
			fmt.Printf("当前已存在 firstBlock.Hash %x\n", lastHash)
		}
		return nil
	})

	bc = BlockChain{
		db:       db,
		lastHash: lastHash,
	}
	return &bc
}

// 追加区块链数据
func (b *BlockChain) AddBlock(tsx []*Transaction) {
	// 这里验证交易
	for _, tx := range tsx {
		//fmt.Printf("矿机池:%d，结果%v\n",i,tx)
		//fmt.Printf("矿机池:%d，结果%v\n",i,b.ValidTransaction(*tx))
		if !b.ValidTransaction(*tx) {
			panic("矿机池发现无效交易，验证失败")
		}
	}
	//初始化块信息
	var newBlock Block
	//从数据库读出最后一个区块key的hash信息，作为自己的preHash值
	var lastHash []byte
	b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_BUCKET))
		lastHash = bucket.Get([]byte(LASH_HASH))
		if lastHash == nil {
			panic("请先初始化区块链")
		}
		return nil
	})
	//保存data信息
	//newBlock.Data = []byte(data)

	// 保存交易体
	newBlock.Transactions = tsx
	//保存preHash值
	newBlock.PreHash = lastHash

	//通过POW结构体计算得到自身的hash值和nonce值
	newBlock.Hash, newBlock.Nonce = NewProofOfWork(&newBlock).Run()
	fmt.Printf("newBlock %x\n", newBlock)

	//保存到数据库中
	b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_BUCKET))
		//保存当前区块
		bucket.Put(newBlock.Hash, newBlock.Serialize())
		//保存当前区块的hash值作为最后一个lastHash
		bucket.Put([]byte(LASH_HASH), newBlock.Hash)

		//同时保存最后一个区块的数据
		b.lastHash = newBlock.Hash
		fmt.Printf("b.lastHash %x\n", b.lastHash)
		return nil
	})

	//// 初始化块信息
	//var newBlock Block
	//// 得到上一个区块的Hash的值，作为自己的preHash值
	//var lenBlocks int = len(b.Blocks)
	//lastBlock := b.Blocks[lenBlocks-1]
	//newBlock.PreHash = lastBlock.Hash
	//// 保存data信息
	//newBlock.Data = []byte(data)
	//// 计算自己的Hash值
	////newBlock.SetHash()
	//
	////修改为通过系统计算得到比系统预设hash小的hash值，刚好可以得到nonce随机值
	//newBlock.Hash, newBlock.Nonce = NewProofOfWork(&newBlock).Run()
	//// 追加到区块链里面
	//b.Blocks = append(b.Blocks, &newBlock)
}

//返回指定地址能够⽀配的utxo所在的交易的集合
func (bc *BlockChain) FindUTXOs(publicKeyHash []byte) []TXOutput {
	var UTXO []TXOutput
	txs := bc.FindTransactions(publicKeyHash)
	for _, tx := range txs {
		for _, output := range tx.TXOutputs {
			//if output.ScriptPubKey == address {
			if bytes.Equal(output.ScriptPubKeyHash, publicKeyHash) {
				UTXO = append(UTXO, output)
			}
		}

	}
	return UTXO
}

// 返回指定地址能够⽀配的utxo所在的交易的集合
func (bc *BlockChain) FindNeedUTXOs(fromPublicKeyHash []byte, amount float64) (map[string][]uint64, float64) {
	utxos := make(map[string][]uint64)
	var calc float64

	// 找到需要的utxo即可，逻辑接近于找全部交易
	txs := bc.FindTransactions(fromPublicKeyHash)

	for _, tx := range txs {
		for currentTXOutputIndex, output := range tx.TXOutputs {
			//拿着自己的公钥hash，看是否等于自己将要消耗的output的公钥hash（锁定脚本）
			if bytes.Equal(output.ScriptPubKeyHash, fromPublicKeyHash) {
				if calc < amount {
					utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], uint64(currentTXOutputIndex))
					calc += output.Value
					if calc > amount {
						return utxos, calc
					}
				}
			}
		}
	}
	if calc < amount {
		fmt.Printf("不满足转账金额,当前总额：%f， 目标金额: %f\n", calc, amount)
	}

	return utxos, calc
}

//使用FindTransaction优化找utxo的过程
func (bc *BlockChain) FindTransactions(fromPublicKeyHash []byte) []*Transaction {
	var txs []*Transaction

	//======

	//我们定义一个map来保存消费过的output，key是这个output的交易id，value是这个交易中索引的数组
	//map[交易id][]int64
	spentOutputs := make(map[string][]int64)

	//创建迭代器
	it := NewBlockChainIterator(bc)
	//1 遍历区块
	for {
		block := it.Next()
		//2 遍历交易
		for _, tx := range block.Transactions {
			//fmt.Printf("current 交易id: %x\n", tx.TXID)

		OUTPUT:
			//3 遍历output，找到和自己相关的utxo（在添加utxo之前检查一下是否已经消耗过了）
			// currentOutputIndex 代表来到了上一个个区块
			for currentOutputIndex, output := range tx.TXOutputs {
				//当前交易的output
				//在这类做一个过滤，把所有已经消耗过的output和即将添加的output进行对比
				//如果相同，那么就不添加，否则才会添加
				//如果当前交易的id存在于我们已经标识的map，那么说明这个交易里面有消耗过的output

				//spentOutputs是上一次循环根据上一次循环中的交易体关联查询到的这次要用的outputs
				//这样子spentOutputs和currentOutputIndex就处于一个交易体里面了
				if spentOutputs[string(tx.TXID)] != nil { // 代表的是当前交易区块是否含有消耗过的output，第一层拦截，如果没有的话，就无须做任何的判断
					for _, preBlockOutputIndex := range spentOutputs[string(tx.TXID)] {
						if preBlockOutputIndex == int64(currentOutputIndex) {
							continue OUTPUT
						}
					}
				}

				//正常流程
				//如果这个output和我们的目标的地址相同，那么就追加到output数组里面
				//但是优化了一下代码，返回全部的含有带有当前地址的交易体返回出去，然后让调用的函数自行处理，外边的函数需要过滤地址
				// if output.ScriptPubKey == address {
				if bytes.Equal(output.ScriptPubKeyHash, fromPublicKeyHash) {
					//UTXO = append(UTXO, output)
					txs = append(txs, tx)
				}
			}

			//如果当前交易是挖矿交易的话，那么不做遍历，直接跳过
			//假如是挖矿，那么重新循环
			//如果不挖矿，那么不循环
			//================
			if tx.IsCoinBase() {
				continue
			}
			//==================
			//否则不是挖矿交易的话，那么就是处理每个交易里面的input
			//遍历input，找到自己各个input对应的上一个区块的output，把消耗过标记出来
			//判断一下每一个input的对应的上一个区块output的sig签名是不是同一个人，是的话才是这个人消耗过的output
			for _, input := range tx.TXInputs {
				//if input.ScriptSig == address {
				if bytes.Equal(HashPublicKey(input.PublicKey), fromPublicKeyHash) {
					//大坑，tagOutList生成一个临时变量，并不是对spentOutputs的引用，
					//造成了上面的TXOutputs处理逻辑每次都是循环的空数据，无法做到有效的过滤已经消耗过的output

					//tagOutputList := spentOutputs[string(input.TXID)]
					////input.VoutIndex是当前input对应的上一个区块的output所在的 交易体的对应的output索引
					//tagOutputList = append(tagOutputList, input.VoutIndex)
					spentOutputs[string(input.TXID)] = append(spentOutputs[string(input.TXID)], input.VoutIndex)
				}
			}
		}
		if len(block.PreHash) == 0 {
			fmt.Println("当前区块链遍历完成")
			break
		}
	}

	//	====
	return txs
}

//签名函数由Transaction提供，但是签名动作由blockChain来实现，因为我们要遍历账本，
//在SignTransaction内部再调用Sign函数

//遍历账本，通过遍历账本的input，依次找到input所属的output所在的交易里面
func (bc *BlockChain) SignTransaction(currentTX *Transaction, privateKey ecdsa.PrivateKey) {
	//每一个交易创建后的最后签名
	everyInputPrevTXs := make(map[string]Transaction)

	//找到全部的交易，当前新建的交易体里面有多少个input，就遍历几遍，从而找到全部的input所引用的交易体
	for _, input := range currentTX.TXInputs {
		tmpPrevTX, err := bc.FindTransactionByInputTXid(input.TXID)
		if err != nil {
			panic(err)
		}
		everyInputPrevTXs[string(tmpPrevTX.TXID)] = tmpPrevTX
	}
	//进行签名
	currentTX.Sign(privateKey, everyInputPrevTXs)
}

//寻找input引用的output，通过output找到对应的交易体，
//在一个当前交易里面，可能有2个input引用的都是同一个交易的里面的output，所以可以用map结构存储所引用的交易体
func (bc *BlockChain) FindTransactionByInputTXid(inputTXid []byte) (Transaction, error) {
	//使用迭代器找到input所属的output的所在的交易体
	it := NewBlockChainIterator(bc)
	for {
		block := it.Next()

		//这里是通过id找交易体
		for _, tx := range block.Transactions {
			//有些是没有交易id的，比如挖矿的交易，所以无法比较
			if bytes.Equal(tx.TXID, inputTXid) {
				return *tx, nil
			}
		}
		if len(block.PreHash) == 0 {
			break
		}
	}
	return Transaction{}, errors.New("无法通过给定的input的TXid找到对应的交易")
}

// 验证每一笔交易
//找到所有引用的交易
//1. 查找当前交易的input，有几个input就遍历几次
//2. 找到input所对应的交易体
//3. 形成map结构存储所有的input交易
//最后调用验证函数验证每一个交易
//根据id查找交易本身，需要遍历整个区块链
func (bc *BlockChain) ValidTransaction(tx Transaction) bool {
	if tx.IsCoinBase() {
		return true
	}
	everyInputPrevTXs := make(map[string]Transaction)
	for _, input := range tx.TXInputs {
		inputTXOne, err := bc.FindTransactionByInputTXid(input.TXID)
		if err != nil {
			panic(err)
		}
		everyInputPrevTXs[string(input.TXID)] = inputTXOne
	}
	return tx.Valid(everyInputPrevTXs)
}
