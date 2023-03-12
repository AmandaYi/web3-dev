package main

import "fmt"

func main() {
	// 统计一个英文字符串，每个字母出现的次数
	const D string = "You use more to indicate that there is a greater amount of something than before or than average, or than something else. You can use 'a little', 'a lot', 'a bit', 'far', and 'much' in front of more ."
	m := make(map[string]int)

	// 取D的长度，依次追加进m中，如果m中已有了这个key就把这个key的值加一，如果没有这个key就新建这个key，并把值设置为1
	dLen := len(D)

	for i := 0; i < dLen; i++ {
		k := string(D[i])
		if k == " " {
			continue
		}
		v, isV := m[k]
		if isV {
			m[k] = v + 1
		} else {
			m[k] = 1
		}
	}
	for k, v := range m {
		fmt.Printf("%s:%d,", k, v)
	}
}
