package main

import "fmt"

func main() {
	// 不惯defer在什么位置，均为最后执行，不过如果是对应panic的话，一定要先声明defer进行异常捕获
	// 如果一个函数中有多个defer的话，会采用压栈的方式处理defer顺序，也就是后进先出，最后进入的最先执行，符合函数的压栈模型
	TestDefer()
	/*
		TestDefer
		defer3
		defer2
		defer1
	*/
}

func TestDefer() {
	defer func() {
		fmt.Println("defer1")
	}()
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	fmt.Println("TestDefer")
}
