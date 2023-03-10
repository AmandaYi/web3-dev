package main

import "fmt"

func main() {
	slice := make([]string, 2, 4)
	slice = []string{"1", "2"}

	slice = append(slice, "3", "4")
	fmt.Println("len:", len(slice), "cap:", cap(slice)) // len: 4 cap: 4
	slice = append(slice, "5", "6")
	fmt.Println("len:", len(slice), "cap:", cap(slice)) // len: 6 cap: 8

	slice2 := slice[0:2:2]
	// append追加新切片数据，并不会为原来的切片追加数据, 并且，超过了容量2，因此与原有的容量扩容方式一样，也是翻倍扩容
	slice2 = append(slice2, "7")
	fmt.Println("slice", slice, "len:", len(slice), "cap:", cap(slice))     // slice [1 2 3 4 5 6] len: 6 cap: 8
	fmt.Println("slice2", slice2, "len:", len(slice2), "cap:", cap(slice2)) //slice2 [1 2 7] len: 3 cap: 4,
}
