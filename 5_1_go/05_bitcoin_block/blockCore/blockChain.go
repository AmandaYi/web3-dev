package blockCore

import (
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
		fmt.Printf("lastHash %x\n", lastHash)
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
	//fmt.Println("--------BlockChain--STR--------")
	//for _, v := range b.Blocks {
	//	v.PrintAll()
	//}
	//fmt.Println("--------BlockChain--END--------")
}

// 初始化区块链 - 链表
func NewBlockChain() *BlockChain {
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
			firstBlock := GenesisBlock(genesisInfo, []byte{})
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
func (b *BlockChain) AddBlock(data string) {
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
	newBlock.Data = []byte(data)
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
