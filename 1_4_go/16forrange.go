package main

import "fmt"

func main() {
	//range是一个函数
	var nums []int = []int{1, 2, 3}
	for k, v := range nums {
		nums[k] = k
		fmt.Println(v)
	}
	fmt.Println(nums)
	//range 返回的v是修改的时候不会影响原数组或者切片或者map或者原结构体,这里的v是一份拷贝

	//需要强调的是range函数返回的永远是浅拷贝
	type student struct {
		name string
	}
	m := make(map[int]student)

	m[1] = student{name: "张三"}

	for k, v := range m {
		v.name = "李四"
		fmt.Println(k, v)
	}

}
