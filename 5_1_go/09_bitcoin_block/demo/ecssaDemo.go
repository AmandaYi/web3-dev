package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {
	//1. 使用椭圆曲线算法生成一对：公钥+私钥
	//2. 使用私钥进行签名
	//3. 使用公钥进行验证
	//
	//golang内置的椭圆曲线库支持签名校验，不支持公钥加密，私钥解密

	//1. 使用椭圆曲线算法生成一对：公钥+私钥
	curve := elliptic.P256()
	randNumber := rand.Reader
	fmt.Println("randNumber", randNumber)

	privateKey, err := ecdsa.GenerateKey(curve, randNumber)
	if err != nil {
		panic(err)
	}
	//2. 使用私钥进行签名
	data := "hello world!"
	randNumber1 := rand.Reader
	fmt.Println("randNumber1", randNumber1)
	dataHashed := sha256.Sum256([]byte(data))
	r, s, err := ecdsa.Sign(randNumber1, privateKey, dataHashed[:])
	if err != nil {
		fmt.Println(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Println("要传输的字节流", signature)

	//3. 使用公钥进行验证
	publicKey := privateKey.PublicKey

	//privateKey.Public() ??// 为什么不能使用？我觉得吧返回的是一个公钥地址引用

	r1 := signature[:len(signature)/2]
	s1 := signature[len(signature)/2:]

	r1BigInt := big.Int{}
	r1BigInt.SetBytes(r1)
	s1BigInit := big.Int{}
	s1BigInit.SetBytes(s1)

	result := ecdsa.Verify(&publicKey, dataHashed[:], &r1BigInt, &s1BigInit)
	fmt.Println("验证结果", result)
}
