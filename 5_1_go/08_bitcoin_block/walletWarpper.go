package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)

const wwFileName string = "walletWrapper.dat"

type WalletWrapper struct {
	// key 地址， value 钱包
	WalletMap map[string]*Wallet
}

//1.  从本地加载已有的钱包到内存
//2.  添加新的钱包到内存
//3.  将内存中的钱包保存到本地

func NewWalletWrapper() *WalletWrapper {
	var ww WalletWrapper
	ww.WalletMap = make(map[string]*Wallet)

	//这里初始化的时候，加载一下文件内存到结构体中，并返回
	err := ww.LoadWalletWrapperFromFile()
	if err != nil {
		panic(err)
	}
	//这里返回局部变量的地址，其实是利用了闭包的原理
	return &ww
}

//创建钱包容器,使用容器来操作内部的单个钱包（单个钱包可以理解为银行卡），
func (ww *WalletWrapper) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.GenAddress()
	ww.WalletMap[address] = wallet

	ww.SaveToFile()
	return address
}

// 保存到本地文件中
func (ww *WalletWrapper) SaveToFile() {
	var content bytes.Buffer

	//gob: type not registered for interface: elliptic.p256Curve
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&content)

	err := encoder.Encode(&ww)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(wwFileName, content.Bytes(), 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 读取本地钱包文件
func (ww *WalletWrapper) LoadWalletWrapperFromFile() error {
	_, err := os.Stat(wwFileName)
	// 如果不是一个是否存在的错误，那么就报错
	if os.IsNotExist(err) {
		return err
	}
	content, err := ioutil.ReadFile(wwFileName)
	if err != nil {
		return err
	}

	var fileWalletWrapper WalletWrapper

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewBuffer(content))
	err = decoder.Decode(&fileWalletWrapper)
	if err != nil {
		return err
	}
	// 设置到内存中
	ww.WalletMap = fileWalletWrapper.WalletMap
	return nil
}

// 获取所有的地址
func (ww *WalletWrapper) GetListWallet() []string {
	var listAddress []string
	for addressKey, _ := range ww.WalletMap {
		fmt.Println("addressKey", addressKey)
		//fmt.Println("value",value)
		listAddress = append(listAddress, addressKey)
	}
	return listAddress
}
