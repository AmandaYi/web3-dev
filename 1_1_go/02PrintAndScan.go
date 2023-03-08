package main

import "fmt"

func main() {
	// 输出
	num1, num2, num3 := 10, 20, 30
	fmt.Println("num1 = ", num1, "num2 = ", num2, "num3 = ", num3)
	//fmt.Print("use Print, num2 = ", num2)
	fmt.Printf("use Printf, num3 = %d\n", num3)

	fmt.Print("123\n")

	//输入
	//https://blog.csdn.net/weixin_51261234/article/details/123934925
	//获取变量的内存地址，输入的时候，必须使用取地址&标识变量名
	//1 Scan
	var num4 int
	fmt.Println("使用Scan接受数据，如果要输入两个值，必须输入两个，如果是三个值，就必须是三个。否则他会一直等待，不会执行下面的程序，Scan必须等待所要求的值输入完成才能执行程序")
	fmt.Scan(&num4)
	fmt.Println("接收到的值是", num4)

	//2 Scanf
	var name string
	fmt.Println("使用Scanf接受数据，格式为name:xxx，Scanf规定用户在输入的时候除了输入相关的&+变量还要在前面加上“%相关类型”")
	fmt.Scanf("name:%s", &name)
	fmt.Println("接受到的值是", name)

	//3 Scanln
	var name2 string
	fmt.Println("使用Scanln接受数据，直接输入，遇到回车结束录入, 遇到输入的内容有空格，会自动使用第一组输入的值， Scanln不管输入的数据是否完成，只要回车就直接执行")
	fmt.Scanln(&name2)
	fmt.Println("接受到的值是", name2)
}
