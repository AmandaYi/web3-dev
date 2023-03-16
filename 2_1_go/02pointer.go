package main

import (
	"fmt"
)

func main() {
	//1 指针注意事项
	// 空指针，只声明，不初始化的的指针会出现空指针问题。
	//var a *int
	// 野指针，被一片无效的地址空间初始化。
	//var p *int = 0x11111;

	//在heap上申请一块空间

	var p *int = new(int)
	fmt.Println(*p)

	var s string
	fmt.Printf("%s", s)
	fmt.Printf("%q", s) // "" 专门可以用q打印go语言的字符串

	//2 变量存储
	//左值，右值
	var a int = 1
	var b int = 1
	a = b
	//等号左边的变量代表所指向的内存空间，（写）
	//等号右边的变量代表变量内存空间存储的数据值 （读）

	//函数传参一些类型传引用的原因是，一个变量在heap堆上保存，但是函数在栈帧内部，因为不在一个空间，所以需要传递参数地址
}
