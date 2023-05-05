package blockCore

import (
	"github.com/boltdb/bolt"
)

type BlockChainIterator struct {
	db               *bolt.DB
	currentBlockHash []byte
}

func NewBlockChainIterator(bc *BlockChain) *BlockChainIterator {
	var it BlockChainIterator
	it.db = bc.db
	it.currentBlockHash = bc.lastHash
	return &it
}

func (it *BlockChainIterator) Next() Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_BUCKET))
		tmpBlock := bucket.Get([]byte(it.currentBlockHash))
		block = block.Deserialize(tmpBlock)
		it.currentBlockHash = block.PreHash
		return nil
	})
	return block
}
