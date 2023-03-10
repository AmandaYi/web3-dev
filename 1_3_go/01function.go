package main

import "fmt"

func main() {
	// 定义参数
	TestFuncArgs("Go", "Go")
	// 不定参数
	TestFuncRestArgs(1, 2, 3, 4, 5, 6)

	// 返回一个值
	ReturnOneValue()
	//返回多个值
	ReturnMoreValue()
}

func TestFuncArgs(name string, code string) {

}

func TestFuncRestArgs(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}

	for _, v := range args {
		// fmt.Printf("k = %d, v = %d\n", k, v)
		fmt.Printf("v = %d\n", v)
	}
}

func ReturnOneValue() int {
	return 1
}

func ReturnMoreValue() (int, string) {
	return 1, "Go"
}
