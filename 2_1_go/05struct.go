package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 结构体是一种数据类型，使用的时候，必须用type定义
	var u1 User = User{
		username: "admin",
		password: "123456",
	}

	var u2 User = User{"admin", "123456"}

	var u3 User = User{username: "admin"} // 部分初始化

	//结构体比较
	fmt.Println(u1 == u2)

	//结构体传参，使用的值拷贝 -- 几乎不用，消耗内存大，一般定义成指针形式，同时指针类型还能保住操作同一份属性

	//结构体地址就是第一个成员的地址

	//查看结构体大小
	fmt.Println(unsafe.Sizeof(u3))

	//指针变量的定义和初始化
	var u4 *User = &User{username: "admin"}
	var u5 *User = new(User)

	//结构体传参,使用引用类型

	//结构体指针做函数返回值，
	//不能返回局部变量的地址值，-- 局部变量保存在栈帧上，函数调用结束后， 栈帧释放，那么对应的局部变量的地址，不在受系统保护，随时会被GC或者分配给其他程序
}

type User struct {
	username string
	password string
}
