package main

import "fmt"

func main() {
	var age int
	var num, sum int
	fmt.Println(age)
	fmt.Println(num, sum)

	// 初始化
	var count int = 100
	fmt.Println(count)

	var c int
	c = 10
	fmt.Println(c)
	var a, b int = 1, 2
	c = a + b

	fmt.Println(c)

	//交换2个变量
	var x, y int = 20, 30
	var tmp int
	tmp = x
	x = y
	y = tmp
	fmt.Println("交换后的值是", x, y)
	fmt.Println(`tmp的值是`, tmp)

	//自动推导类型
	autoDerivation := 10
	fmt.Println(autoDerivation)
	autoDerivation1, autoDerivation2, autoDerivation3 := 10, 11, 13
	fmt.Println(autoDerivation1, autoDerivation2, autoDerivation3)
	//交换2个变量
	swapVariable1, swapVariable2 := 100, 200
	swapVariable1, swapVariable2 = swapVariable2, swapVariable1
	fmt.Println("使用多个变量赋值交换2个变量", swapVariable1, swapVariable2)
}
