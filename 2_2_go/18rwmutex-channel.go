package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 读写锁和信道一起使用，造成了隐形死锁

	//下面的代码是错误的，无法对数据进行准确处理
	var ch chan int = make(chan int, 5)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		go readInt2(ch, i)
	}
	for i := 0; i < 5; i++ {
		go writeInt2(ch, i)
	}

	time.Sleep(time.Second * 2)
}

func readInt2(rCh <-chan int, index int) {
	for {
		n := <-rCh
		fmt.Printf("第%d读RRR Go程读出了数据%d\n", index, n)
		time.Sleep(time.Second)
	}
}
func writeInt2(wCh chan<- int, index int) {
	for {
		n := rand.Intn(1000)
		wCh <- n
		fmt.Printf("第%d写WWWW Go程写入了数据%d\n", index, n)
		time.Sleep(time.Second)
	}
}
