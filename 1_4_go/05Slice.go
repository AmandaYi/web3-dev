package main

import "fmt"

func main() {
	// 切片相对于数组，长度不是固定的，可以追加元素，在追加时容量是可以不断增大的，可以理解为动态数组，但是他不是数组

	var slice1 []string
	slice1 = append(slice1, "Hello")
	slice1 = append(slice1, "Go")

	fmt.Printf("%s\n", slice1[0])

	slice2 := []string{"Hi", "World"}
	fmt.Printf("%s\n", slice2)

	// 使用make创建切片，make(type, len, cap) ,len代表初始化多少空间，cap代表整个切片多大【后面可以翻倍扩容】
	slice3 := make([]int, 5, 10)
	fmt.Printf("slice3:%d, len:%d, cap:%d\n", slice3, len(slice3), cap(slice3)) // slice3:[0 0 0 0 0], len:5, cap:10

	slice4 := make([]string, 10, 10)
	fmt.Printf("slice3:%s, len:%d, cap:%d\n", slice4, len(slice4), cap(slice4)) // slice4:[          ], len:5, cap:10

	// 切片的截取 s[low:high:max]
	// 注意新切片是指向原有的切片的，因此对新切片的修改，影响原来的切片
	slice5 := []int{1, 2, 3, 4, 5, 6}
	slice6 := slice5[0:3:5]                                                     // 这里的容量就是第三个值5减去第一个值0得到
	fmt.Printf("slice6:%d, len:%d, cap:%d\n", slice6, len(slice6), cap(slice6)) // slice6:[1 2 3], len:3, cap:5
}
