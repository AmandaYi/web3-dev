package main

import "fmt"

func main() {
	// 数组
	var Numbers [5]int = [5]int{1, 2, 3, 4, 5}
	for _, v := range Numbers {
		fmt.Println(v)
	}
	var Names [2]string = [2]string{"n1", "n2"}
	for _, n := range Names {
		fmt.Println(n)
	}
	// 部分赋值
	var Numbers2 [5]int = [5]int{1, 2, 3}
	for _, v2 := range Numbers2 {
		fmt.Println(v2)
	}

	//通过初始化的值，进行计算长度，一般不常用
	Number3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("长度是%d", len(Number3))
}
