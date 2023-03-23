package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	timer := time.NewTimer(time.Second * 5)
	//可以重置计时器，
	timer.Reset(time.Second * 2)
	//可以停止计时器，如果停止了计时器，那么下边的读取操作就会卡死了，并报错fatal error: all goroutines are asleep - deadlock!
	// timer.Stop()
	<-timer.C // 这一步在阻塞，因为timer.NewTimer一个写chan一直没有执行，那么这里的读chan就会阻塞，定时到了，系统会往time.NewTimer中的C信道写入数据后，这里才能读，才不阻塞

	fmt.Println(time.Now())

	//定时任务的三种方式
	//1. 使用封装好的sleep，底层还是使用了NewTimer
	time.Sleep(time.Second)
	//2. 计时器底层
	time.NewTimer(time.Second)
	//3. 再次封装的NewTimer，返回的是time.NewTimer的C信道
	<-time.After(time.Second)

	//让main函数利用信道阻塞原理等着AfterFunc的回调函数
	waitMainTest := make(chan bool)
	// 多少时间后做什么事情，接受一个函数, 返回Time对象
	_ = time.AfterFunc(time.Second, func() {
		fmt.Println("time.AfterFunc")
		waitMainTest <- true
	})

	<-waitMainTest
	fmt.Println("waitMainTest")
}
