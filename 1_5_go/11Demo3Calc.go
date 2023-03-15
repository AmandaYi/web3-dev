package main

import "fmt"

func main() {
	var f CalcFactory
	fmt.Println(f.Create("+", 1, 2))
}

//多态主要解决的是需求改变后，需要修改对象创建代码的问题
//使用工厂类，内部调用多态函数，即可解决此问题

type BaseCalc struct {
	num1 int
	num2 int
}

type AddCalc struct {
	BaseCalc
}

type ReduceCalc struct {
	BaseCalc
}

func (ac *AddCalc) GetResult() int {
	return ac.num1 + ac.num2
}

func (rc *ReduceCalc) GetResult() int {
	return rc.num1 - rc.num2
}

type CalcER interface {
	GetResult() int
}

func CalcGetResult(h CalcER) int {
	return h.GetResult()
}

type CalcFactory struct {
}

func (cc *CalcFactory) Create(op string, num1, num2 int) int {

	switch op {
	case "+":
		var addCalc AddCalc = AddCalc{BaseCalc{
			num1: num1,
			num2: num2,
		}}
		return CalcGetResult(&addCalc)
	case "-":
		var reduceCalc ReduceCalc = ReduceCalc{BaseCalc{
			num1: num1,
			num2: num2,
		}}
		return CalcGetResult(&reduceCalc)
	default:
		return 0
	}

}
