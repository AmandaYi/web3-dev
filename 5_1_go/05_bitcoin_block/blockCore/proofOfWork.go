package blockCore

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	// 区块数据
	block *Block
	// 目标值，计算的系统指定的,要小于这个值
	target *big.Int
}

//添加创建PoW的函数，根据计算流程，返回的应该是比系统值小的hash和试算数nonce

func NewProofOfWork(b *Block) *ProofOfWork {

	pow := ProofOfWork{
		block: b,
	}
	/*
		//系统目标值
		targetString := "0000f00000000000000000000000000000000000000000000000000000000000"
		//进行复制
		var tmpTarget big.Int
		tmpTarget.SetString(targetString, 16)
		pow.target = &tmpTarget
	*/

	targetLocal := big.NewInt(1)
	targetLocal.Lsh(targetLocal, 242)
	pow.target = targetLocal
	return &pow

}

//工作量证明处理函数
func (pow *ProofOfWork) Run() (hash []byte, nonce uint64) {
	//得到当前区块
	//b := pow.blockCore
	//初始化一个nonce的值
	nonce = 0
	hash = []byte{}
	var tmpHash [32]byte
	var tmpBigInt big.Int
	//初始化一个sha256，然后进行对比目标值即可，找到了就返回hash
	for {
		//genesisInfoByte := [][]byte{
		//	b.PreHash, b.Data, uint64ToByte(b.Version), b.MerKleRoot, uint64ToByte(b.TimeStamp), uint64ToByte(b.Difficulty), uint64ToByte(nonce),
		//}
		//tmp := bytes.Join(genesisInfoByte, []byte{})
		tmpHash = sha256.Sum256(pow.prepareData(nonce))
		tmpBigInt.SetBytes(tmpHash[:])
		//比较找打一个一个比系统目标值小的hash即可
		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		//func (x *Int) Cmp(y *Int) (r int) {
		if tmpBigInt.Cmp(pow.target) == -1 {
			hash = tmpHash[:]
			fmt.Printf("挖矿成功！hash : %x, nonce : %d\n", hash, nonce)
			break
		} else {
			nonce++
		}
	}
	return hash, nonce
}

//进行数据字段拼接，返回一个可以进行hash计算的值
//prepareData
func (pow *ProofOfWork) prepareData(nonce uint64) []byte {
	b := pow.block
	genesisInfoByte := [][]byte{
		b.PreHash, b.Data, uint64ToByte(b.Version), b.MerKleRoot, uint64ToByte(b.TimeStamp), uint64ToByte(b.Difficulty), uint64ToByte(nonce),
	}
	return bytes.Join(genesisInfoByte, []byte{})
}
