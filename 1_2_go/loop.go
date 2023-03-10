package main

import "fmt"

func main() {
	//TestForNormal()

	//TestForLoop()

	// 计算营业额
	TestBilling()

	// 水仙花
	TestNarcissus()

	// 乘法口诀
	TestMulitplyTable()

	TestMulitplyTable2()
}
func TestForNormal() {
	var i int
	var count int = 10000
	for i = 0; i < count; i++ {
		fmt.Printf("%d媳妇，我爱你! ", i+1)
		if i%100 == 0 {
			fmt.Printf("\n")
		}
	}
	defer func() {
		fmt.Printf("\n")
	}()
}

func TestForLoop() {
	for {
		fmt.Println("请输入我爱你")
		var input string
		fmt.Scanf("%s", &input)
		if input == "我爱你" {
			fmt.Println("我也爱你")
			break
		} else {
			fmt.Println("呜呜，你变心了")
		}
	}
}

func TestBilling() {
	var initBilling float64 = 80000
	var count = 0
	for {
		var tmp float64 = initBilling * 1.25
		count++
		if tmp >= 200000 {
			initBilling = tmp
			break
		} else {
			initBilling = tmp
		}

	}
	fmt.Printf("%d\n", count+2006)
	fmt.Println(initBilling)
}
func TestNarcissus() {
	var start, end int = 100, 999
	var index int = start

	for {
		if index > end {
			break
		}

		var hundreds int = int(index / 100) // 百位

		var decade int = index / 10 % 10 // 十位

		var onePlace int = index % 10 // 个位

		if onePlace*onePlace*onePlace+decade*decade*decade+hundreds*hundreds*hundreds == index {
			fmt.Println(index)
		}
		index++
	}
}

func TestMulitplyTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d x %d = %d,", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}

func TestMulitplyTable2() {
	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%d x %d = %d,", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}
