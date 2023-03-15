package main

import (
	"fmt"
	"strings"
)

func main() {
	Test()
}

func Test() {
	var s string = "123456789abcdefgABCDEFG!@#$%^&*()_+"
	fmt.Println(strings.Contains(s, "1"))          // 是否包含
	fmt.Println(strings.Join([]string{s, s}, ",")) // 字符串连接起来
	fmt.Println(strings.Index(s, "1"))             // 查找字符串所在的位置
	fmt.Println(strings.Repeat(s, 1))              // 查找重复多少次的字符
	fmt.Println(strings.Replace(s, "a", "666", 2)) // 替换字符, 最后的n是代表找几次
	fmt.Println(strings.Split(s, "a"))             // 以什么字符进行，分割字符得到切片
}
