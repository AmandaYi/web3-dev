package main

import "fmt"

func main()  {
	//常量定义：程序运行期间，不可改变的值， 推荐全大写定义常量名
	const VERSION string = "2.0"

	//打印常量地址, 无法打印常量地址
	// fmt.Printf("%p", &VERSION) // cannot take the address of VERSION

	fmt.Printf("%s", VERSION)

}
