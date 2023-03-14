package main

import "fmt"

type Add struct {
	Object
}
type Reduce struct {
	Object
}

type Object struct {
	num1 int
	num2 int
}

func (p *Add) GetResult(a int, b int) int {
	p.num1 = a
	p.num2 = b
	return p.num1 + p.num2
}

func (p *Reduce) GetResult(a int, b int) int {
	p.num1 = a
	p.num2 = b
	return p.num1 - p.num2
}

func main() {

	var add Add = Add{}
	fmt.Println(add.GetResult(1, 2))
	var reduce Reduce = Reduce{}
	fmt.Println(reduce.GetResult(1, 2))
}
