package main

import "fmt"

func main() {
	//字典，映射，key为基本类型，不能是引用类型
	//var m map[int]string // 无法使用，不能存数据，必须进行初始化才可以
	//m[1] = "1" // panic: assignment to entry in nil map
	//fmt.Println(m)

	//m2 := map[int]string{}
	m3 := make(map[int]string, 5) //但是不能指定cap
	fmt.Println(len(m3))          // 0

	//判断是否有某个key
	//v, ok := m3[1]

	//delete删除map中某个key，即使不存在的key也不会报错
	delete(m3, 1234)
}
