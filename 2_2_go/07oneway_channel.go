package main

import (
	"fmt"
	"time"
)

func main() {
	// 单向通道channel

	//定义

	var c1 chan<- int = make(chan<- int, 5) // 只写通道
	var c2 <-chan int = make(<-chan int, 5) // 只读通道
	var c3 chan int = make(chan int, 2)     // 双向通道

	//转换，
	//	双向通道可以转换为单向通道，传递双向通道赋值，默认传过去的是引用
	c1 = c3
	c2 = c3
	fmt.Println(c1, c2)
	//单向通道不可以强转为双向通道的
	// c3 = c1 // cannot use c1 (type chan<- int) as type chan int in assignment

	//一个生产消费小例子
	//go ProductInt(c3)
	go ProductInt(c1) // 传入只写就行，因为前面c3把地址给了c1，所以c1相当是只有写功能的c3，其实还是c3

	//主go程进行消费
	//ConsumerInt(c3)
	ConsumerInt(c2) // 传入只读就行，因为前面c3把地址给了c2，所以c2相当是只有写功能的c3，其实还是c3
}

//生产一些数据，传入只写chan
func ProductInt(c chan<- int) {
	for i := 0; ; i++ {
		c <- i
		//time.Sleep(time.Millisecond  * 100)
		time.Sleep(time.Second)
	}
}

//消费一些数据，传入只读chan
func ConsumerInt(c <-chan int) {
	for {
		fmt.Printf("%d ", <-c)
	}
}
