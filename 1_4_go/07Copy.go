package main

import "fmt"

func main() {
	slice1 := make([]int, 2, 2)
	slice1 = []int{1, 2}
	slice2 := make([]int, 4, 4)
	fmt.Println("slice2:", slice2, "len:", len(slice2), cap(slice2)) // slice2: [0 0 0 0] len: 4 4
	// 只要初始化了，就会重新计算len和cap，也可以理解为[]int代表了一个新切片赋值过来了，如果用append的话，那么就不会改变len和cap
	slice2 = []int{10, 20, 30}
	fmt.Println("slice2:", slice2, "len:", len(slice2), cap(slice2)) // slice2: [10 20 30] len: 3 3
	copy(slice2, slice1)
	fmt.Println("slice1:", slice1, "len:", len(slice1), cap(slice1)) // slice1: [1 2] len: 2 2
	fmt.Println("slice2:", slice2, "len:", len(slice2), cap(slice2)) // slice2: [1 2 30] len: 3 3

	slice3 := make([]int, 4, 4)
	fmt.Println("slice3:", slice3, "len:", len(slice3), cap(slice3)) // slice3: [0 0 0 0] len: 4 4
	//如果用append的话，那么就不会改变len和cap,下面这行代码，len+1，cap*2
	slice3 = append(slice3, 1)
	fmt.Println("slice3:", slice3, "len:", len(slice3), cap(slice3)) // slice3: [0 0 0 0 1] len: 5 8
}
