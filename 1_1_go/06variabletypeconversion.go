package main

import "fmt"

func main() {
	// 强制类型转换，用于将一种数据类型的变量转换为另外一种类型的变量
	var num float64 = 3.14926
	fmt.Printf("%d\n", int(num));

	var num1 int = 10;
	fmt.Printf("%.2f\n", float64(num1))

	//在类型转换时候，一定建议把低类型转换为高类型，否则的话，把高类型转换成低类型，可能出现精度丢失和数据溢出。
	//比如
	var num2 int = 200
	fmt.Printf("%d\n", int8(num2)) // -56
}
