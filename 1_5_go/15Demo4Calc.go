package main

import "fmt"

func main() {
	// 这里使用类型断言和空接口实现数据校验，改造了计算器程序
	var f D4CalcFactory
	op := f.Create("+")
	if op.SetData(1, 2) {
		fmt.Println(D4CalcGetResult(op))
	}
}

//多态主要解决的是需求改变后，需要修改对象创建代码的问题
//使用工厂类，内部调用多态函数，即可解决此问题
//定义基础类，比如打印机
type D4BaseCalc struct {
	num1 int
	num2 int
}

//定义高级类，比如彩色打印机
type D4AddCalc struct {
	D4BaseCalc
}

//定义高级类，比如黑色打印机
type D4ReduceCalc struct {
	D4BaseCalc
}

//定义接口
type D4CalcER interface {
	GetResult() int
	SetData(data ...interface{}) bool // 使用断言和空接口进行判断数据类型
}

//定义接口多态方法
func D4CalcGetResult(h D4CalcER) int {
	return h.GetResult()
}

//实现接口方法,用来匹配接口
func (ac *D4AddCalc) GetResult() int {
	return ac.num1 + ac.num2
}

//实现接口方法,用来匹配接口
func (ac *D4AddCalc) SetData(data ...interface{}) bool {
	var result bool = true
	if len(data) > 2 {
		result = false
		return result
	}
	{
		v, ok := data[0].(int)
		if !ok {
			fmt.Println("类型错误，只能为int")
			result = false
			return result
		}
		ac.num1 = v
	}
	{
		v, ok := data[1].(int)
		if !ok {
			fmt.Println("类型错误，只能为int")
			result = false
			return result
		}
		ac.num2 = v
	}
	return result
}

//实现接口方法,用来匹配接口
func (rc *D4ReduceCalc) GetResult() int {
	return rc.num1 - rc.num2
}

//实现接口方法,用来匹配接口
func (rc *D4ReduceCalc) SetData(data ...interface{}) bool {
	var result bool = true
	if len(data) > 2 {
		result = false
		return result
	}
	{
		v, ok := data[0].(int)
		if !ok {
			fmt.Println("类型错误，只能为int")
			result = false
			return result
		}
		rc.num1 = v
	}
	{
		v, ok := data[1].(int)
		if !ok {
			fmt.Println("类型错误，只能为int")
			result = false
			return result
		}
		rc.num2 = v
	}
	return result
}

//创建工厂类,用于调用创建高级类,调用多态方法
type D4CalcFactory struct {
}

func (cc *D4CalcFactory) Create(op string) D4CalcER {

	switch op {
	case "+":
		return new(D4AddCalc)
	case "-":
		return new(D4ReduceCalc)
	default:
		return nil
	}

}
