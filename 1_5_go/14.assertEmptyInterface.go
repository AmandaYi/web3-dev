package main

import "fmt"

func main() {
	//通过类型断言判断接口内部的类型
	//语法： value, ok := m.(T)
	//m：表示空接口类型变量
	//T：是断言的变量
	//value：变量m中的值
	//ok：布尔类型的量，如果断言成功为true，否则为false

	var i interface{}
	i = 1
	value, ok := i.(int) // 这里用value, ok := i.(interface) 是永远成功的，因为任何类型都实现了空接口
	if ok {
		fmt.Printf("%d", value)
	} else {
		fmt.Printf("类型推断错误")
	}
}
