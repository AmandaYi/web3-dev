package main

import "fmt"

func main() {
	var number [3][3]string = [3][3]string{{"1-1", "1-2", "1-3"}, {"2-1", "2-2", "2-3"}, {"3-1", "3-2", "3-3"}}

	for _, v := range number {
		for _, vInner := range v {
			fmt.Printf("数据分别是%s", vInner)
		}
	}

}
