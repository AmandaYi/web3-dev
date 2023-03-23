package main

import (
	"fmt"
	"sync"
	"time"
)

//在上层应用里面，全部是建议锁，建议锁是操作系统提供的，建议编程使用。
//强制锁是操作系统底层自己使用的

var mutex sync.Mutex // 创建一个互斥量，新的互斥量状态为0

func main() {
	go user1()
	go user2()
	time.Sleep(time.Second * 30)
}

func printer(s string) {
	mutex.Lock()
	for _, v := range s {
		fmt.Printf("%c", v)     // 争夺stdout屏幕
		time.Sleep(time.Second) // CPU不断的在换执行人
	}
	mutex.Unlock()
}

func user1() {
	printer("你好")
}
func user2() {
	printer("123456")
}
