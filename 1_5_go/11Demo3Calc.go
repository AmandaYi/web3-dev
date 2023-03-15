package main

func main() {

}

//多态主要解决的是需求改变后，需要修改对象创建代码的问题
//使用工厂类，内部调用多态函数，即可解决此问题

type CaleCommon struct {
	num1 int
	num2 int
}

type ADDCalc struct {
	CaleCommon
}

type ReduceCalc struct {
	CaleCommon
}

func (ac *ADDCalc) GetResult() int {
	return ac.num1 + ac.num2
}

func (rc *ReduceCalc) GetResult() int {
	return rc.num1 - rc.num2
}

type CCalcFactory struct {
}

func (cc *CCalcFactory) CreateOperator(op string, num1, num2 int) int {

	switch op {
	case "+":
		var addCalc ADDCalc = ADDCalc{CaleCommon{
			num1: num1,
			num2: num2,
		}}
		return CalcGetResult(&addCalc)
	case "-":
		var reduceCalc ReduceCalc = ReduceCalc{CaleCommon{
			num1: num1,
			num2: num2,
		}}
		return CalcGetResult(&reduceCalc)
	default:
		return 0
	}

}

type CCalcEr interface {
	GetResult() int
}

func CalcGetResult(h CCalcEr) int {
	return h.GetResult()
}
