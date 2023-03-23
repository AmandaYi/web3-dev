package main

import (
	"fmt"
	"time"
)

func main() {
	// 周期计时器
	//Ticker 跟 Timer 的不同之处，就在于 Ticker 时间达到后不需要人为调用 Reset 方法，会自动续期。
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Println(<-ticker.C)
	}
}
