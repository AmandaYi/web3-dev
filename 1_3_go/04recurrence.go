package main

func main() {
	TestRecurrence()
	//计算n!阶乘
	println(Factorial(5))
}
func TestRecurrence() {

}
func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
