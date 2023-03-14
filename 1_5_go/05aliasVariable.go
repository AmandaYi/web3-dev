package main

type Int int

func (a Int) TestInt(b Int) Int {
	return a + b
}

// 这是不可以的，基本变量不能作为一个对象接受者，需要用type给基本变量起别名
//func (a int) TestGolangInt() {
//
//}
