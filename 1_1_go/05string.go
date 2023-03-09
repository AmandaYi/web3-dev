package main

import "fmt"

func main() {
	// 字符串
	//用双引号包起来
	//字符串结束标志是\0
	//使用len()函数得到字符串的长度, 不会计算最后的\0这个长度，只会计算\0之前的个数
	//Golang语言中一个汉字占用三个字节
	var name string = "你好"
	fmt.Println(len(name)) // 6
}
