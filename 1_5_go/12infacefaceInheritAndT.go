package main

import "fmt"

func main() {
	var order Order
	var orderER OrderER = &order
	orderER.CreateOrder()
	orderER.Log() // 可以调用接口自己继承的方法，不过前提是结构体实现了，不过如果结构体没有实现的话，也就不能把结构体赋值给该接口了

	var baseOrderER BaseOrderER = orderER
	baseOrderER.Log() // 接口的转换，接口的低等级的可以赋值给高等级的，反过来是不可以的
}

type BaseOrderER interface {
	Log()
}
type OrderER interface {
	BaseOrderER
	CreateOrder()
}

type Order struct {
}

func (order *Order) CreateOrder() {
	fmt.Println("创建订单")
}

func (order *Order) Log() {
	fmt.Println("订单日志")
}
