package main

import "fmt"

func main()  {

	var name string
	var age int

	fmt.Println("请输入姓名:")

	_, _ = fmt.Scan(&name)

	fmt.Println("请输入年龄:")

	_, _ = fmt.Scanln(&age)

	fmt.Printf("您好:%s您的年龄是%d", name, age)
}
