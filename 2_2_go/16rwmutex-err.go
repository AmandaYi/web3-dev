package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex sync.RWMutex

func main() {
	// 读写锁和信道一起使用，造成了隐形死锁,不推荐锁和信道一起使用
	rand.Seed(time.Now().UnixNano())
	var ch chan int = make(chan int)
	for i := 0; i < 5; i++ {
		go readInt(ch, i)
	}
	for i := 0; i < 5; i++ {
		go writeInt(ch, i)
	}

	time.Sleep(time.Second * 10)
}

func readInt(rCh <-chan int, index int) {
	for {
		rwMutex.RLock()
		n := <-rCh
		fmt.Printf("第%d读RRR Go程读出了数据%d\n", index, n)
		rwMutex.RUnlock()
	}
}
func writeInt(wCh chan<- int, index int) {
	for {
		n := rand.Intn(1000)
		rwMutex.Lock()
		wCh <- n
		fmt.Printf("第%d写WWWW Go程写入了数据%d\n", index, n)
		rwMutex.Unlock()
		time.Sleep(time.Second)
	}
}
