package main

import "fmt"

func main() {
	slice := make([]int, 2, 2)
	// 切片的函数参数传递模式是引用传递，与数组的值传递是不同的
	InitSlice(slice)
	fmt.Println("slice=", slice, "len=", len(slice), "cap", cap(slice))
}

func InitSlice(s []int) {
	sLen := len(s)
	sCap := cap(s)

	for i := 0; i < sLen; i++ {
		s[i] = i
	}
	fmt.Println("slice=", s, "len=", sLen, "cap", sCap)
}
