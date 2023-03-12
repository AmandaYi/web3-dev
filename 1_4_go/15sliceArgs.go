package main

import "fmt"

func main() {
	//把切片作为参数可以修改切片中的元素，原因分析如下
	//slice底层是结构体，传参只会浅拷贝，但是slice结构体里面的array里面的指针也同时拷贝过来了，但是却没有深度拷贝开辟新空间，所以造成了引用了切片
	/*type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}
	*/
	//array作为底层数组，只是把array引用地址传递了，因此造成了在函数里面修改切片，会影响到函数外边的切片内容

	//可以通过append方法对切片进行动态扩容
	//append在追加数据的时候，会先检测len的长度，
	//必须先检测slice底层数组是否有足够的容量来保存新添加的元素。如果有足够空间的话，直接扩展slice（依然在原有的底层数组之上），将新添加的y元素复制到新扩展的空间，并返回slice。
	//因此，输入的x和输出的z共享相同的底层数组。
	//
	//如果没有足够的增长空间的话，appendInt函数则会先分配一个足够大的slice用于保存新的结果，先将输入的x复制到新的空间，然后添加y元素。结果z和输入的x引用的将是不同的底层数组。
	//
	s := make([]int, 5, 20)
	ModifySlice(s)
	fmt.Println("原始数组", s)
}

func ModifySlice(s []int) {
	//sLen := len(s)
	//for i := 0; i < sLen; i++ {
	//	s = append(s, i)
	//}
	//s[0] = 123
	//fmt.Println("函数内部数组", s)

	//假如说还能往里面放入一个元素
	s = append(s, 123)
	fmt.Println("函数内部数组", s)
}
