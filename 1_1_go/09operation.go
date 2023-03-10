package main

import "fmt"

func main() {
	// + - * / % ++ --

	// go语言没有前++或者前--， 只有后++，后--，永远是一个整体

	// 关于计算类的，推荐都声明为float64
	//计算圆的面积和周长
	TestAreaOfCircleAndCircumferenceOfCircle()
	// 总分和平均分
	CalculateTheTotalScoreAndAverageScore()

	// 购买T恤，裤子求和，并舍去小数位
	BugShirtTrousers()

	// = += -= *= /= %=

	// == != < > <= >=

	// 逻辑运算符
	// ! && ||
	// && 优先级高于||

	// 单目运算符，就是只需要一个操作数即可，比如++, --, !, &[取地址]
	// 双目运算符，就是必须两个操作参与运算
	// +, -, *, /, <= ,>=

	// 测试闰年
	TestLeapYear()
}

//计算圆的面积和周长
func TestAreaOfCircleAndCircumferenceOfCircle() {
	const PI float64 = 3.14
	var r float64 = 10
	var Area, Circumference float64 = 0, 0
	Area = r * r * PI
	Circumference = 2 * PI * r
	fmt.Printf("面积是%.2f, 周长是%.2f", Area, Circumference)
	defer func() {
		fmt.Printf("\n")
	}()
}

// 总分和平均分
func CalculateTheTotalScoreAndAverageScore() {
	const CHINESE, MATH, ENGLISH float64 = 90, 89, 69

	var sum, average float64 = 0, 0
	sum = CHINESE + MATH + ENGLISH
	average = sum / 3

	fmt.Printf("总分是%0.f, 平均分是%0.f", sum, average)
	defer func() {
		fmt.Printf("\n")
	}()
}

// 购买T恤，裤子求和，并舍去小数位
func BugShirtTrousers() {
	const SHIRTPRICE, TROUSERSPRICE float64 = 35, 120
	const SHIRTCOUNT, TROUSERS float64 = 3, 2
	const DISCOUNT float64 = 8.8
	var sum float64 = SHIRTPRICE*SHIRTCOUNT + TROUSERSPRICE*TROUSERS
	fmt.Printf("应该付%f", sum)
	fmt.Printf("舍去小数部分为%.0f元", sum)
	defer func() {
		fmt.Printf("\n")
	}()
}

// 判断闰年
func TestLeapYear() {
	var year int
	//是否是闰年
	var isLeap bool = false

	fmt.Printf("请输入一个年份")

	_, _ = fmt.Scanln(&year)

	// 符合两个条件之一
	if (year%200 == 0) || (year%4 == 0 && (year%100 != 0)) {
		isLeap = true
	}
	if isLeap {
		fmt.Printf("%d是闰年", year)
	} else {
		fmt.Printf("%d不是闰年", year)
	}
}
