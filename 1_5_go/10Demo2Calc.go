package main

import "fmt"

func main() {
	var cInt CInt
	fmt.Println(CAdd(&cInt, 1, 2))
	fmt.Println(CReduce(&cInt, 2, 2))
}

//用多态模拟实现，实现一个加减计算器

type CalcEr interface {
	Add(num1, num2 int) int
	Reduce(num1, num2 int) int
}

func CAdd(h CalcEr, num1, num2 int) int {
	return h.Add(num1, num2)
}
func CReduce(h CalcEr, num1, num2 int) int {
	return h.Reduce(num1, num2)
}

type CInt int

func (i *CInt) Add(num1, num2 int) int {
	return num1 + num2
}
func (i *CInt) Reduce(num1, num2 int) int {
	return num1 + num2
}
