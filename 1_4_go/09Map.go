package main

import "fmt"

func main() {
	var map1 map[string]string = map[string]string{"name": "小明"}
	fmt.Println(map1)
	map2 := map[string]string{}
	map2["name"] = "小明"
	fmt.Println(map2)
	map3 := make(map[string]string)
	map3["name"] = "小明"
	fmt.Println(map3)
}
