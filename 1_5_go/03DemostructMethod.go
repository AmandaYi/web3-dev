package main

import "fmt"

func main() {
	var s1 Student = Student{
		Person: Person{"小明", 1, 18},
		Source: Source{100, 100, 100},
	}
	s1.Hello() // 虽然定义的是指针接受者，但是值类型依旧可以使用，但是会隐式传入指针值，相当于 (&s1).Hello
	(&s1).SourceSum()
	(&s1).Average()

	// 1 指针接受者（s *Student），在使用的时候，如果其中一个方法改变了属性的值，那么该接受者的另一个方法也会得到修改后的属性
	//带来的效果：任何对指针接收者的修改会体现到 原调用者。

	// 2 非指针接受者(s Student), 在使用的时候，不会互相应用，即使使用&s作为接受者，只是一份复制的值的地址而已，并不会影响到另一个方法里面访问的对应的接受者

	//什么时候使用指针接收者
	//需要对接受者的变更能体现到原调用者
	//当struct占用很大内存，最好使用指针接受者，否则每次调用接受者函数 都会形成struct的大副本

	//http://www.manongjc.com/detail/63-hvsmwfbnhfqvudu.html
}

type Person struct {
	name   string
	gender uint // 1 男 2 女
	age    uint
}

type Source struct {
	Chinese uint
	Math    uint
	English uint
}

type Student struct {
	Person
	Source
}

func (s *Student) Hello() {
	var gender uint = s.gender
	var genderName string
	if gender == 1 {
		genderName = "男同学"
	} else if gender == 2 {
		genderName = "女同学"
	} else {
		genderName = "保密"
	}
	fmt.Printf("大家好，我是%s,今年%d岁了，我性别是%s\n", s.name, s.age, genderName)
}
func (s *Student) SourceSum() {
	var sum uint = s.Chinese + s.English + s.Math
	fmt.Printf("我总分是%d\n", sum)
}
func (s *Student) Average() {
	var average uint = (s.Chinese + s.English + s.Math) / 3
	fmt.Printf("我的平均分是%d\n", average)
}
