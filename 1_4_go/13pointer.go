package main

import "fmt"

// 指针是一个变量，存起来的是变量的地址

func main() {
	var a int
	var p *int = &a
	fmt.Printf("%T\n", &a) // *int
	fmt.Println(p)         // 0xc00000a0a8

	var p2 *int
	fmt.Println(p2) // 空指针 <nil>

	var p3 *int
	fmt.Println(&p3) // 0xc0000d8020

	var p4 *int = new(int)
	*p4 = 100
	fmt.Println("p4指向的地址是", p4, "该地址保存的存储空间的值是", *p4) // p4指向的地址是 0xc000102058 该地址保存的存储空间的值是 100

	//数组指针
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var p5 *[10]int
	p5 = &nums
	fmt.Println(*p5, (*p5)[0]) // []的优先级高于*p, 也可以写成p[0], 是语言优化
	UpdateArr(p5)
	fmt.Println(p5[0])

	//指针数组
	//一个数组里面每一项都是指针，就叫做指针数组
	var arrP [5]*int // 这里的*int代表每一项都是一个指针
	arrP = [5]*int{new(int), &a, new(int), new(int)}
	fmt.Println(arrP) // [0xc00000a0f0 0xc00000a0a8 0xc00000a0f8 0xc00000a100 <nil>]
	//怎么取出某一个数
	fmt.Println(*(arrP[1])) // 这里需要括起来arrP[1] 这样子是先找到数组，然后用*去找这个内容项对应的具体值，跟上面的数组指针是不同的。

	//通过数组指针和指针数组，体现了*预算符的用法。它的作用就是寻找某个地址对应的地址单元的具体内容值

	//切片指针
	var sliceList []int = []int{1, 2, 3}
	var sliceP *[]int = &sliceList
	fmt.Println(*sliceP, sliceP, (*sliceP)[0])

	//指针切片
	var sliceA int = 1
	var sliceB int = 2
	var sliceList2 []*int = []*int{&sliceA, &sliceB}
	fmt.Println(sliceList2, *(sliceList2[0]))

	//结构体指针
	type Student struct {
		name string
	}
	student1 := Student{"张三"}
	var studentP *Student = &student1
	fmt.Println("结构体用法", (*studentP).name, studentP.name) // .的优先级是比*高的，因此需要把*studentP括起来, 不加*也可以用简化的写法

	// 多级指针
	var abc int = 1
	var pAbc *int = &abc
	var qAbc **int = &pAbc
	var mAbc ***int = &qAbc
	fmt.Println("通过多级指针取到abc变量存的值", abc, *pAbc, **qAbc, ***mAbc)

	//总结，*取值符号,&是取地址符号
}

func UpdateArr(p *[10]int) {
	p[0] = 10
}
