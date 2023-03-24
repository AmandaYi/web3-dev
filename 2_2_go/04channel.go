package main

import (
	"fmt"
	"time"
)

var channel = make(chan int)

func main() {
	//程序打开后，自动打开三个文件叫做stdout,stdin,stderr,程序关闭后，操作系统自动关闭这三个文件

	//channel 基本原理是队列， 是一种数据类型

	//定义
	_ = make(chan int)    // 无缓冲的channel
	_ = make(chan int, 0) // 无缓冲的channel
	_ = make(chan int, 1) // 有缓冲的channel, 在存满之前，不会阻塞

	ch := make(chan int, 10)
	fmt.Println("ch len = ", len(ch)) // 返回内部有多少已保存的数据，也就是大家常说的剩余未使用的数据个数 ch len =  0
	fmt.Println("ch cap", cap(ch))    // 返回总的缓冲区大小 ch cap 10

	//channel有两个端，一端：写端（传入端） chan <- xxx
	//另一端：读端（传出端） xxx <- chan
	//要求，读端和写端必须同时满足条件，才能在chan上进行数据流动，否则，则阻塞

	go User1()
	go User2()
	time.Sleep(time.Second * 30)
}

func Printer(s string) {

	for _, v := range s {
		fmt.Printf("%c", v)     // 争夺stdout屏幕
		time.Sleep(time.Second) // CPU不断的在换执行人
	}
}

func User1() {
	Printer("你好")
	channel <- 1 // 写数据
}
func User2() {
	<-channel // 读数据，在内部无数据的时候，读不到数据的时候，此行代码暂停，进入堵塞
	Printer("123456")
}
