package main

import "math/big"

type ProofOfWork struct {
	// 区块数据
	block *Block
	// 目标值，计算的系统指定的,要小于这个值
	target *big.Int
}

//添加创建PoW的函数，根据计算流程，返回的应该是比系统值小的hash和试算数nonce

func NewProofOfWork() (hash []byte, nonce uint64) {

}
