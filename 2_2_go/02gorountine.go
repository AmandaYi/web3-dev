package main

import (
	"fmt"
	"time"
)

func main() {
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
	}
}
