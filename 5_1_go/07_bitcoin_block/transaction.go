package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"

	"fmt"
)

const reward = 12.5

// 交易输入TXInput
//指明交易发起⼈可⽀付资⾦的来源，包含：
//引⽤utxo所在交易的ID
//所消费utxo在output中的索引
//解锁脚本
type TXInput struct {
	// 引用utxo集合所在交易的id
	TXID []byte
	// 引用output的索引值
	VoutIndex int64
	// 解锁脚本
	ScriptSig string
}

//交易输出（TXOutput）
//包含资⾦接收⽅的相关信息,包含：
//接收⾦额
//锁定脚本
//==易错点：经常把Value写成⼩写字⺟开头的==，这样会⽆法写⼊数据库，切记！
type TXOutput struct {
	// 接受的金额
	Value float64
	// 锁定脚本
	ScriptPubKey string
}

//区块中的交易结构，该交易结构自我hash生成MerKleRoot值
type Transaction struct {
	//本次交易ID
	TXID []byte
	//交易的输入，可以是多个
	TXInputs []TXInput
	//交易的输出，可以是多个
	TXOutputs []TXOutput
}

//设置交易的id的hash值
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(tx)
	data := buffer.Bytes()
	idHash := sha256.Sum256(data)
	tx.TXID = idHash[:]
}

// 创建一个第一条交易
// coinbase总是新区块的第⼀条交易，这条交易中只有⼀个输出，即对矿⼯的奖励，没有输⼊。
func NewCoinBaseTX(address string, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("reward %s %f\n", address, reward)
	}
	var tx Transaction
	//唯一输入
	//⽐特币系统，对于这个input的id填0，对索引填0xffff，data由矿⼯填写，⼀般填所在矿池的名字
	var input TXInput = TXInput{
		TXID:      []byte{},
		VoutIndex: -1,
		ScriptSig: data,
	}

	//唯一输出
	var output TXOutput = TXOutput{
		Value:        reward,
		ScriptPubKey: address,
	}

	tx.TXInputs = []TXInput{input}
	tx.TXOutputs = []TXOutput{output}

	tx.SetHash()
	return &tx
}

// 解锁脚本，解锁脚本是为了找到自己的余额，也就是utxo
//解锁脚本是检验input是否可以使⽤由某个地址锁定的utxo，所以对于解锁脚本来说，是外部提供锁定
//信息，我去检查⼀下能否解开它。
//我们没有涉及到真实的⾮对称加密，所以使⽤字符串来代替加密和签名数据。即使⽤地址进⾏加密，
//同时使⽤地址当做签名，通过对⽐字符串来确定utxo能否解开。
func (input *TXInput) CanUnlockUTXOWith(unlockData string) bool {
	// ScriptSign 是签名信息，因为前面直接用付款人的地址作为了签名，因此这里先拿来进行比较
	return input.ScriptSig == unlockData
}

//锁定脚本
//同样的，锁定脚本是⽤于指定⽐特币的新主⼈，在创建output时创建。对于这个output来说，它应该是
//⼀直在等待⼀个签名的到来，检查这个签名能否解开⾃⼰锁定的⽐特币。
func (output *TXOutput) CanBeUnlockedWith(unlockData string) bool {
	return output.ScriptPubKey == unlockData
}

// 判断是不是挖矿交易
func (tx *Transaction) IsCoinBase() bool {
	// 条件是，只有一个input，且这一个input的交易id为nil，且这一个input的 引用output的索引值（VoutIndex）为-1
	var result bool = false
	if len(tx.TXInputs) == 1 {
		//if bytes.Equal(tx.TXInputs[0].TXID, []byte{}) && tx.TXInputs[0].VoutIndex == -1 {
		//	result = true
		//}
		//因为序列化会把[]byte压缩为nil，造成无法对比问题，所以这里使用长度对比
		if tx.TXInputs[0].VoutIndex == -1 {
			result = true
		}
	}
	return result
}

//创建一个普通的转账交易
// 需要有来源人，目的人，金额，哪条区块链
//过程是
//找到需要的合理的utxo
//创建一个交易的输入input
//创建交易输出

//如果自己的utxo总额少于了需要付的交易金额，那么不去处理，并给出提示
//如果自己的utxo总额超过了需要付的交易金额，那么就给自己也创造一个output给自己找零
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	var tx Transaction
	var inputs []TXInput
	var outputs []TXOutput

	//找到需要的合理的utxos map[string][]uint64
	utxos, utxoSumAmount := bc.FindNeedUTXOs(from, amount)
	//如果自己的utxo的总额比本次需要付的交易金额要小，就不往下走了
	if utxoSumAmount < amount {
		fmt.Printf("余额不足，当前余额%f,付账总额为%f\n", utxoSumAmount, amount)
		return nil
	}

	//创建一个交易的输入input
	//utxos里面存的是，可以用来直接支付给下一个区块的output
	for TXid, canUseOutputsIndex := range utxos {
		for _, VoutIndex := range canUseOutputsIndex {
			input := TXInput{
				TXID:      []byte(TXid),
				VoutIndex: int64(VoutIndex),
				ScriptSig: from,
			}
			inputs = append(inputs, input)
		}
	}

	//创建交易的输出
	output := TXOutput{
		Value:        amount,
		ScriptPubKey: to,
	}
	outputs = append(outputs, output)
	//如果自己的utxo超过了需要支付的金额，那么就把多出去的零钱转给自己，也就是多写一个output
	if utxoSumAmount > amount {
		output := TXOutput{
			Value:        utxoSumAmount - amount,
			ScriptPubKey: from, // 转给自己,锁定脚本也用自己的公钥
		}
		outputs = append(outputs, output)
	}

	tx = Transaction{
		TXID:      nil,
		TXInputs:  inputs,
		TXOutputs: outputs,
	}
	tx.SetHash()
	return &tx
}
