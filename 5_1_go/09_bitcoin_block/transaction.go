package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"math/big"
	"strings"

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
	//ScriptSig string

	//签名, 真正的签名，由r，s拼接的字节流，其中r和s是长度相等的
	Signature []byte

	// 公钥，这里存储的公钥，并不是最原始的公钥，而是存储极坐标X和Y的字节流，在校验时候可以拆开得到公钥
	// 这里是公钥，不是哈希，也不是地址，是付款人的公钥
	PublicKey []byte
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
	//ScriptPubKey string

	//收款方的公钥哈希，是公钥的哈希，不是地址
	ScriptPubKeyHash []byte
}

// 在调用输出的时候，一定是知道了收款人的地址，因此可以拿收款人的地址还原出来收款人的公钥HASH
func (output *TXOutput) Lock(address string) {
	hash := GetPublicKeyHashFromAddress(address)
	output.ScriptPubKeyHash = hash
}

// 使用方法创建交易输出，提高封装性
//创建交易的时候，能拿到金额和收款人的地址
func NewTxOutput(value float64, address string) TXOutput {
	output := TXOutput{
		Value:            value,
		ScriptPubKeyHash: nil,
	}
	// 设置锁定脚本,也就是设置收款人的公钥HASH的值
	output.Lock(address)
	return output
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
	//这个hash并没有对全部的数据进行什么处理，只是对数据结构进行了签名,以后可能会优化
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
	input := TXInput{
		TXID:      []byte{},
		VoutIndex: -1,
		Signature: nil,
		PublicKey: []byte(data), // 这里是挖矿交易，因此没有所谓的公钥，一般让挖矿者设置为自己的矿池的名字
	}

	//唯一输出
	//var output TXOutput = TXOutput{
	//	Value:        reward,
	//	//ScriptPubKey: address,
	//}
	output := NewTxOutput(reward, address)

	//对于挖矿交易来说，只有一个input和一output
	tx.TXInputs = []TXInput{input}
	tx.TXOutputs = []TXOutput{output}

	tx.SetHash()
	return &tx
}

/**
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
*/
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
	//我们创建新的交易一定是要使用钱包里面的公钥私钥，所以整个步骤如下：
	//1.  打开钱包，根据创建人的address找到对应的钱包（银行卡）
	//2.  查找可用的utxo，注意此时传递的不再是地址，而是地址的公钥哈希：pubKeyHash
	//3.  创建输入
	//4.  创建输出（付款，找零）
	//5.  使用私钥对交易进行签名
	//通过付款人地址->找到钱包->找到公钥+私钥
	ww := NewWalletWrapper()
	wallet := ww.WalletMap[from]
	if wallet == nil {
		fmt.Println("没有找到付款人地址对应的钱包，交易失败！")
		return nil
	}
	privateKey := wallet.PrivateKey
	publicKey := wallet.PublicKey

	//计算公钥的hash值
	publicKeyHash := HashPublicKey(publicKey)

	var tx Transaction
	var inputs []TXInput
	var outputs []TXOutput

	//找到需要的合理的utxos map[string][]uint64
	//utxos, utxoSumAmount := bc.FindNeedUTXOs(from, amount)

	utxos, utxoSumAmount := bc.FindNeedUTXOs(publicKeyHash, amount)

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
				//ScriptSig: from,
				Signature: nil,
				PublicKey: publicKey,
			}
			inputs = append(inputs, input)
		}
	}

	//创建交易的输出
	//output := TXOutput{
	//	Value:        amount,
	//	ScriptPubKey: to,
	//}
	output := NewTxOutput(amount, to)

	outputs = append(outputs, output)
	//如果自己的utxo超过了需要支付的金额，那么就把多出去的零钱转给自己，也就是多写一个output
	if utxoSumAmount > amount {
		//output := TXOutput{
		//	Value:        utxoSumAmount - amount,
		//	ScriptPubKey: from, // 转给自己,锁定脚本也用自己的公钥
		//}
		output := NewTxOutput(utxoSumAmount-amount, from)
		outputs = append(outputs, output)
	}

	tx = Transaction{
		TXID:      nil,
		TXInputs:  inputs,
		TXOutputs: outputs,
	}
	tx.SetHash()

	// 创建交易的时候进行签名
	bc.SignTransaction(&tx, *privateKey)
	return &tx
}

// 签名核心
//前提，要用私钥，同时根据交易的input里面的交易id，找到所有的引用的交易体

// 实现逻辑是

//1. 对TXInputs里面的每一项input都要做签名
//2. 逻辑是：首先拿到自己的私钥和即将签名的交易(tx)
//3. 创建一个当前交易的副本，同时把副本的input的Signature和PublicKey设置为nil
//4. 循环input所引用的交易体，得到input所引用的output（也就是能够支配的utxo）的公钥hash，把值付给当前交易的PublicKey
//5. 然后对当前交易整体做hash计算，可以直接调用当前交易生成TX.TXID的SetHash函数，反正是副本怎么弄都行
//6. 还原当前input的PublicKey的值为nil，以免影响当前交易的其他input,同时把用过的交易体的TXID也还原回去
//7. 把新生成的TX.TXID作为一个临时变量存起来
//8. 进行椭圆ECDSA签名，得到R和S，保存到当前交易的input的签名里面
/**
@param privateKey
@param everyInputPrevTX 每一个input引用的output的交易体，这是output所在的交易体
*/
func (tx *Transaction) Sign(privateKey ecdsa.PrivateKey, everyInputPrevTXs map[string]Transaction) {
	//这里是签名逻辑
	if tx.IsCoinBase() {
		return
	}
	copyCurrentTX := tx.CopyTransactionBelongSign()

	for i, input := range copyCurrentTX.TXInputs {
		inputPrevTXOne := everyInputPrevTXs[string(input.TXID)]
		if len(inputPrevTXOne.TXID) == 0 {
			panic("引用的交易无效")
		}
		//4. 循环input所引用的交易体，得到input所引用的output（也就是能够支配的utxo）的公钥hash，把值付给当前交易的PublicKey
		//input.VoutIndex存的是当前input所引用的交易的第几个output（utxo）
		//要是不嫌麻烦，可以再次遍历一遍
		copyCurrentTX.TXInputs[i].PublicKey = inputPrevTXOne.TXOutputs[input.VoutIndex].ScriptPubKeyHash
		//然后对当前交易整体做hash计算，可以直接调用当前交易生成TX.TXID的SetHash函数，反正是副本怎么弄都行
		copyCurrentTX.SetHash()
		//还原当前input的PublicKey的值为nil，以免影响当前交易的其他input,同时把用过的交易体的TXID也还原回去
		copyCurrentTX.TXInputs[i].PublicKey = nil
		//把新生成的TX.TXID作为一个临时变量存起来
		tmpSignatureData := copyCurrentTX.TXID
		//还原交易体的TXID
		copyCurrentTX.TXID = tx.TXID

		//进行椭圆ECDSA签名，得到R和S，保存到当前交易的input的签名里面
		r, s, err := ecdsa.Sign(rand.Reader, &privateKey, tmpSignatureData)
		if err != nil {
			panic(err)
		}
		tx.TXInputs[i].Signature = append(r.Bytes(), s.Bytes()...)
	}

}

//创建副本的函数（同时把副本的input的Signature和PublicKey设置为nil）
func (tx *Transaction) CopyTransactionBelongSign() Transaction {
	var inputs []TXInput
	var outputs []TXOutput
	for _, input := range tx.TXInputs {
		inputs = append(inputs, TXInput{
			TXID:      input.TXID,
			VoutIndex: input.VoutIndex,
			Signature: nil,
			PublicKey: nil,
		})
	}
	for _, output := range tx.TXOutputs {
		outputs = append(outputs, TXOutput{
			Value:            output.Value,
			ScriptPubKeyHash: output.ScriptPubKeyHash,
		})
	}
	//for i, _ := range srcTransaction.TXInputs {
	//	srcTransaction.TXInputs[i].Signature = nil
	//	srcTransaction.TXInputs[i].PublicKey = nil
	//}
	return Transaction{
		TXID:      tx.TXID,
		TXInputs:  inputs,
		TXOutputs: outputs,
	}
}

//校验核心验证每一个交易
//由于交易中已经存储了 数字签名 和 公钥 ，所以只需要将引用的交易传递进来，为了获取引用输出的公钥哈希
//分析校验：
//所需要的数据：公钥，数据(txCopy，生成哈希), 签名
//我们要对每一个签名过得input进行校验
//1. 得到想要签名的数据，使用拷贝函数，同时把拷贝的副本的，得到签名r和s的逻辑和签名是一样的，只是这里不进行后续的操作了
//详细见Sign函数，再得到r和s之后，后面才会不一样
//2. 因为前面input的PublicKey的值从nil->自己的所属的output的PublicKeyHash->nil,相当于最后还是nil，
//所以这里从最原理的交易的input里面重新赋值一份，是为了拿到XY极坐标
//3.通过极坐标XY还原出来公钥,通过input的Signature拆分出来R和S

//4. 进行验证
func (tx *Transaction) Valid(everyInputPrevTXs map[string]Transaction) bool {
	if tx.IsCoinBase() {
		return true
	}
	//1. 得到想要签名的数据，使用拷贝函数，同时把拷贝的副本的，得到签名r和s的逻辑和签名是一样的，只是这里不进行后续的操作了
	//详细见Sign函数，再得到r和s之后，后面才会不一样
	copyCurrentTX := tx.CopyTransactionBelongSign()
	for i, input := range copyCurrentTX.TXInputs {
		inputPrevTXOne := everyInputPrevTXs[string(input.TXID)]
		if len(inputPrevTXOne.TXID) == 0 {
			panic("引用的交易无效")
		}
		//=====
		//4. 循环input所引用的交易体，得到input所引用的output（也就是能够支配的utxo）的公钥hash，把值付给当前交易的PublicKey
		//input.VoutIndex存的是当前input所引用的交易的第几个output（utxo）
		//要是不嫌麻烦，可以再次遍历一遍
		copyCurrentTX.TXInputs[i].PublicKey = inputPrevTXOne.TXOutputs[input.VoutIndex].ScriptPubKeyHash
		//然后对当前交易整体做hash计算，可以直接调用当前交易生成TX.TXID的SetHash函数，反正是副本怎么弄都行
		copyCurrentTX.SetHash()
		//还原当前input的PublicKey的值为nil，以免影响当前交易的其他input,同时把用过的交易体的TXID也还原回去
		copyCurrentTX.TXInputs[i].PublicKey = nil
		//把新生成的TX.TXID作为一个临时变量存起来
		tmpSignature := copyCurrentTX.TXID
		//还原交易体的TXID
		copyCurrentTX.TXID = tx.TXID
		//====
		//拿到三个数据
		//1. input存储的公钥XY
		//inputPublicKey := input.PublicKey // 这个input并不是正确的input
		inputPublicKey := tx.TXInputs[i].PublicKey
		//2. 本次的签名，重签一边，也就是数据copyCurrentTX.TXID，这个值也就是input的签名信息，只是这里没有赋值给当前的input的Signature
		signatureAgainData := tmpSignature
		//3. input存储的签名
		//inputSignature := input.Signature // 这个input并不是正确的input
		inputSignature := tx.TXInputs[i].Signature

		//3.通过极坐标XY还原出来公钥,通过input的Signature拆分出来R和S
		var X, Y, R, S big.Int = big.Int{}, big.Int{}, big.Int{}, big.Int{}
		X.SetBytes(inputPublicKey[:len(inputPublicKey)/2])
		Y.SetBytes(inputPublicKey[len(inputPublicKey)/2:])

		R.SetBytes(inputSignature[:len(inputSignature)/2])
		S.SetBytes(inputSignature[len(inputSignature)/2:])

		//还原公钥
		publicKeyOrigin := ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     &X,
			Y:     &Y,
		}

		//4. 进行验证
		result := ecdsa.Verify(&publicKeyOrigin, signatureAgainData, &R, &S)

		if !result {
			return result
		}
	}
	return true
}

func (tx Transaction) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("--- Transaction %x:", tx.TXID))

	for i, input := range tx.TXInputs {

		lines = append(lines, fmt.Sprintf("     Input %d:", i))
		lines = append(lines, fmt.Sprintf("       TXID:      %x", input.TXID))
		lines = append(lines, fmt.Sprintf("       Out:       %d", input.VoutIndex))
		lines = append(lines, fmt.Sprintf("       Signature: %x", input.Signature))
		lines = append(lines, fmt.Sprintf("       PubKey:    %x", input.PublicKey))
	}

	for i, output := range tx.TXOutputs {
		lines = append(lines, fmt.Sprintf("     Output %d:", i))
		lines = append(lines, fmt.Sprintf("       Value:  %f", output.Value))
		lines = append(lines, fmt.Sprintf("       Script: %x", output.ScriptPubKeyHash))
	}

	return strings.Join(lines, "\n")
}
