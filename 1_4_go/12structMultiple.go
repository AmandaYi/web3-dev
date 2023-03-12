package main

import "fmt"

//多层引用结构体，成员结构体是否传入的是拷贝值还是引用值

//经过测试，即使是结构体内部的结构体，也是整体拷贝到函数参数内部的
type Student struct {
	name   string
	source SourceDescription
}
type SourceDescription struct {
	Chinese int
	English int
	Math    int
}

func main() {
	student1 := Student{name: "张三", source: SourceDescription{Chinese: 100, English: 200, Math: 300}}
	fmt.Println(student1)         // {张三 {100 200 300}}
	ModifySourceChinese(student1) // 修改Chinese的值为120
	fmt.Println(student1)         // {张三 {100 200 300}}
}

func ModifySourceChinese(student Student) {
	student.source.Chinese = 120
}
