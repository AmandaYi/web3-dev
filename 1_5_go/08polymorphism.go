package main

import "fmt"

func main() {
	// 所谓多态，指的是多种表现形式，多态就是同一个接口，使用不同的实例而执行不同操作

	//定义： func 方法名（变量 接口类型名）｛
	// 变量.method
	// ｝

	var c ColorPrinter
	var b BlackPrinter
	WhoDoPrint(&c) // 多态使用
	WhoDoPrint(&b) // 多态使用
}

type Printerer interface {
	DoPrint()
}

type ColorPrinter struct {
}
type BlackPrinter struct {
}

func (c *ColorPrinter) DoPrint() {
	fmt.Println("我是彩色打印机，我打印彩色")
}
func (b *BlackPrinter) DoPrint() {
	fmt.Println("我是黑白打印机，我打印黑白")
}

//多态
func WhoDoPrint(h Printerer) {
	h.DoPrint()
}
