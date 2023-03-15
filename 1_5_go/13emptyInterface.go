package main

import "fmt"

func main() {
	// 空接口，任何类型都实现了空接口
	// 打印类型是，数值是什么类型就是什么类型
	var i interface{}
	i = 1
	i = "1"
	i = 1.11
	fmt.Printf("%T\n", i) // float64

	var slice []interface{}
	slice = append(slice, 1, "1", 1.11)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%T\n", slice[i])
	}

	fmt.Printf("%T", slice) // []interface {}
}
