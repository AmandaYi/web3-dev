package main

import "fmt"

func main() {
	var map1 map[string]string = map[string]string{"name": "小明"}
	fmt.Println(map1)
	map2 := map[string]string{}
	map2["name"] = "小明"
	fmt.Println(map2)
	map3 := make(map[string]string, 2)
	map3["name"] = "小明"
	map3["name2"] = "小明"
	map3["name3"] = "小明"
	//
	fmt.Println(map3, "长度是", len(map3)) // 这里的len返回的是map中已有的key的个数，而不是初始化的时候的指定的长度

	map3 = map[string]string{"name": "小明2", "school": "大学", "source": "100"}

	// 获取某个值
	fmt.Println("map3中的school是", map3["school"])

	// 判断是否有某个key
	source, isSource := map3["source"]
	if isSource {
		fmt.Println("当前分数为", source)
	} else {
		fmt.Println("木有这个key source")
	}

	//遍历
	for k, v := range map3 {
		fmt.Printf("%s=%s,", k, v)
	}
	fmt.Printf("\n")
	//删除某个键
	delete(map3, "source")
	for k, v := range map3 {
		fmt.Printf("%s=%s,", k, v)
	}
	fmt.Printf("\n")

	TestArgsMap(map3)   //  school=大学,name=小明2,
	TestModifyMap(map3) // 修改一下
	TestArgsMap(map3)   // name=小明2school=大学source=101
	// 证明map传参是引用传递
}
func TestArgsMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%s=%s", k, v)
	}
	fmt.Printf("\n")

}
func TestModifyMap(m map[string]string) {
	m["source"] = "101"
}
