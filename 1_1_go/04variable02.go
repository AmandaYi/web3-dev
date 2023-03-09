package main

import "fmt"

func main() {
	// 保留关键字如下
	/*
			break, default, func, interface, select,
			case, defer, go, map, struct,
			chan, else, goto, package, switch,
			const, fallthrough, if, range, type,
			continue, for, import, return, var
			还有很多，
		    true, false, iota, nil,

			int, int8, int16, int32, int64,
			unit, uint8, uint16,uint32, uint64,uintptr(uintptr 实际上就是一个 uint 用来表示地址)
			float32,float64,complex128, complex64,
			bool, byte(byte 等同于int8，常用来处理ascii字符), rune(rune 等同于int32,常用来处理unicode或utf-8字符), string, error,
			make, len, cap(cap主要是为了让slice提供可变长度,反射方式读取，其实len也是反射), new, append, copy, close, delete,
			complex(Golang没有complex类型，只有complex64和complex128两种类型,这里的complex是函数,用于从指定的实部和虚部构建复数), real(获取虚数的实部), imag(获取实数的虚部),
			panic(抛出异常信息，如果没有对异常处理，程序会终止), recover(在defer后置函数中，终止一个信道的错误产生过程，回复代码正常运行，必须要在panic之前声明，最好写在函数入口的第一行，可以获取panic的错误信息用来传递)

	*/
	// 整型, 有符号整型(正整数，负整数，0， 用int类声明)， 无符号整型(正整数，0， 用uint类声明)
	//有符号整数的取值范围，在32位系统里面是-21亿到+21亿，正的最后一位要-1，在64位系统里面是太大了，记不住9开头的19位数字，负19位，正19位，正的最后一位要-1
	//无符号证书的取值范围，在32位系统里是0到42亿，在64位系统里面，18后面19位，跟上面的有符号的2位一样
	var number1 int = -1
	fmt.Println(number1)
	var number2 uint = 1
	fmt.Println(number2)
	// 浮点类型
	// 包含小数点的类型，
	//float32精确到小数点后7位，float64精确掉小数点后15位
	//由于在Go的各种运算包里面大量使用floate64，因此建议以后都用float64
	var number3 float32 = 1.1
	fmt.Println(number3)
	var number4 float64 = 2.2
	fmt.Printf("%f\n", number4)
	//如果有需要进行四舍五入进行保留数据，那么需要格式化输入即可如下
	var number5 float64 = 1.123456
	fmt.Printf("%.3f\n", number5)
	number6 := 1.111
	//自动推导类型默认是float64
	fmt.Printf("%T\n", number6) // float64

	// 布尔类型, 默认值是false，输出格式符是%t
	var isOK bool = true
	fmt.Printf("%t\n", isOK)

	// 字符类型
	// 用单引号引起来的单子字符
	var c1 byte = 'a'
	fmt.Printf("%c\n", c1)
	fmt.Printf("%d\n", c1) // 97 ASCII

	// 字符串类型

	// 字符串与字符类型区别, 会是那种C语言的最后加上\0的区别吗？

	// 强制类型转换, 隐式转换

	// fmt包的格式化输入和输出

	// 归纳

}
