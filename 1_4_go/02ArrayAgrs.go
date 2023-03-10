package main

import "fmt"

var Numbers [5]int = [5]int{1, 2, 3, 4}

func main() {
	TestArrayArgs(Numbers)
	fmt.Println(Numbers[0])

}

// 数组作为参数拷贝传递,是值传递，不是引用传递
func TestArrayArgs(Numbers [5]int) {
	Numbers[0] = 99
	for i := 0; i < len(Numbers); i++ {
		fmt.Println(Numbers[i])
	}

}

// fmt.Println(equal([5]int{1, 2, 3, 4, 5}, [5]int{1, 2, 3, 4}))
// 比较数组
func equal(array1 [5]int, array2 [5]int) (isResult bool) {

	if len(array1) != len(array2) {
		return isResult
	}

	for i := 0; i < len(array1); i++ {
		if array1[i] == array1[2] {

		} else {
			isResult = false
			break
		}
	}
	return isResult
}
