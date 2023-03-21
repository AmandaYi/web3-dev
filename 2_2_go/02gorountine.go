package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	n := runtime.GOMAXPROCS(1) // 返回值是打印上一次的设置的核心数，如果没有设置，那么会返回程序自动获取的值
	fmt.Println(n)             // 4

	go Login()
	go Register()
	time.Sleep(time.Second * 30)
}

func Login() {
	for i := 0; i < 10; i++ {
		fmt.Printf("用户%d登录中\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func Register() {
	for i := 0; i < 10; i++ {
		fmt.Printf("用户%d注册中\n", i)
		time.Sleep(time.Millisecond * 100)
		if i < 2 {
			runtime.Gosched() // 让出当前go协程，等会再次获取cpu时从当前位置接着执行，时间片轮转调度算法。
		}
		if i == 5 {
			runtime.Goexit() // 退出当前go协程, 结束调用该函数的当前Go协程，Goexit()之前注册的defer都生效。
		}
	}
}
