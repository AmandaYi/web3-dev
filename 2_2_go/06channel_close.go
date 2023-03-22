package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int = make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			time.Sleep(time.Second)
		}
		//如果这里知道不用chan了那就关闭即可
		close(c)
	}()
	time.Sleep(time.Second * 5)
	for v := range c { // 这里会堵塞 fatal error: all goroutines are asleep - deadlock!
		fmt.Printf("%d", v)
	}
}
