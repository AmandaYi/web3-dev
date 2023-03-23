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
	time.Sleep(time.Second * 1)
	go func() {
		//单纯的输出
		for v := range c { //  可以用此种方便遍历信道的值
			fmt.Println("for range 也可以用来拿信道的值", v)
		}
	}()
	for {
		var ok bool
		var n int
		if n, ok = <-c; ok == true {
			fmt.Println("这样子也可以拿信道的值", n)
		}

		if ok == false {

			fmt.Println("关闭了，此时的值是零值", <-c) // 即使关闭了，也是可以读取的chan的，并不会报错，读出的是零值
			break
		}
	}

}
