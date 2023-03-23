package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex1 sync.RWMutex
var currentInt int

func main() {
	// 读写锁和信道一起使用，造成了隐形死锁
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		go readInt1(i)
	}
	for i := 0; i < 5; i++ {
		go writeInt1(i)
	}

	time.Sleep(time.Second * 2)
}

func readInt1(index int) {
	for {
		rwMutex1.RLock()

		fmt.Printf("第%d读RRR Go程读出了数据%d\n", index, currentInt)
		rwMutex1.RUnlock()
		time.Sleep(time.Second)
	}
}
func writeInt1(index int) {
	for {
		n := rand.Intn(1000)
		rwMutex1.Lock()
		currentInt = n
		fmt.Printf("第%d写WWWW Go程写入了数据%d\n", index, currentInt)
		rwMutex1.Unlock()
		time.Sleep(time.Second)
	}
}
