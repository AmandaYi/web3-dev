package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go GoChannelLogin()
	defer func() {
		fmt.Println("main defer")
	}()
	for i := 0; i < 10; i++ {
		fmt.Printf("main %d", i)
		time.Sleep(time.Second)
		if i == 5 {
			// runtime.Goexit() // fatal error: no goroutines (main called runtime.Goexit) - deadlock! 在主go程退出之后，子go程虽然还执行，但是最后通知主go程自己执行完毕时候，会报错
			//Goexit是不可以用在主go程的
		}
	}
}

func GoChannelLogin() {
	defer func() {
		fmt.Println("GoChannelLogin defer")
	}()
	for i := 0; i < 10; i++ {
		fmt.Printf("channel %d", i)
		time.Sleep(time.Second)
		if i == 2 {
			runtime.Goexit() // 正常
		}
	}
}
