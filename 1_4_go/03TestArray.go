package main

import "fmt"

func main() {
	FindMax()
}

//求最大整数
func FindMax() {
	var Numbers [5]int = [5]int{10, 2, 3, 4, 51}
	//设最大的数字是max
	var max int = Numbers[0]
	for i := 0; i < len(Numbers); i++ {
		if max < Numbers[i] {
			max = Numbers[i]
		}
	}
	fmt.Printf("最大的数字是%d", max)
}
