package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

const version string = "1"

type Wallet struct {
	// 私钥
	PrivateKey *ecdsa.PrivateKey

	// 公钥（使用XY极坐标拼接）
	PublicKey []byte
}

//创建钱包结构

func NewWallet() *Wallet {
	//创建椭圆算法模型
	curve := elliptic.P256()
	//生成私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	publicKeyTmp := privateKey.PublicKey
	//使用公钥XY极坐标
	publicKey := append(publicKeyTmp.X.Bytes(), publicKeyTmp.Y.Bytes()...)
	var wallet Wallet = Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
	return &wallet
}

//通过公钥按照约定的算法，生成地址
func (wallet *Wallet) GenAddress() string {
	//1. 使用椭圆算法生成非对称公钥和私钥
	publicKey := wallet.PublicKey
	//2. 拿到公钥做sha256后，再做RIPEMD160得到公钥的hash，长度是20bytes
	tmp := sha256.Sum256(publicKey)
	hash160er := ripemd160.New()
	_, err := hash160er.Write(tmp[:])
	if err != nil {
		fmt.Println(err)
		return ""
	}
	ripemd160Hash := hash160er.Sum(nil)
	//3. 在公钥160hash前面加上一个字节的版本号，得到21bytes的data
	payloadHash := append([]byte(version), ripemd160Hash[:]...)
	//4. 复制21bytes的data，做两次sha256的hash，取前4个字节bytes拼接到21bytes的hash后面即可，得到25bytes的数据
	var copyPayloadHash []byte = make([]byte, len(payloadHash))
	copy(copyPayloadHash, payloadHash)
	tmpCopyHash := sha256.Sum256(copyPayloadHash)
	tmpCopyHash = sha256.Sum256(tmpCopyHash[:])
	payloadHash = append(payloadHash, tmpCopyHash[:4]...)
	//5.把25bytes的数据做base58编码得到25位的比特币地址
	return base58.Encode(payloadHash)
}
