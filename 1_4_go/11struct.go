package main

import "fmt"

func main() {

	// 结构体是一组相同类型或者不同类型构成的数据集合

	s1 := Source{
		Chinese: 100,
		English: 101,
		Math:    102,
	}

	fmt.Println(s1.Chinese)

	s2 := Source{
		Chinese: 100,
		English: 110,
		Math:    100,
	}
	// 结构体是值传递，因为底层是数组
	TestArgsStruct(s2)

	fmt.Println(s2)
	//数组结构体
	var souceList1 [2]Source = [2]Source{Source{1, 2, 3}}
	souceList1[0] = Source{1, 2, 3}
	souceList1[1] = Source{4, 5, 6}
	fmt.Println(souceList1)
	// 切片结构体
	var sourceList2 []Source = []Source{Source{Chinese: 1, English: 2, Math: 3}}
	sourceList2 = append(sourceList2, Source{100, 200, 300})
	fmt.Println(sourceList2[0].Chinese)

	//Map结构体
	var m map[string]Source = map[string]Source{"source": {100, 200, 300}}
	fmt.Println(m["source"].Chinese)

}

//结构体，万物均为对象，struct是最顶级的对象，给最顶级的对象用type重新定义一个新的名字，叫做Source，也就是产生了一个Source这样的变量类型
type Source struct {
	Chinese int
	English int
	Math    int
}

func TestArgsStruct(s Source) {
	s.Chinese = 200
}
