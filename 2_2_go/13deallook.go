package main

import "fmt"

func main() {
	//死锁不是锁，是一种现象

	// 1 单Go自己死锁
	//每一个channel应该在两个以上的Go程中出现，才不会造成单Go程死锁
	c1 := make(chan int)
	c1 <- 123 // 这里写入后，写入阻塞，后面就没有执行机会
	num1 := <-c1
	fmt.Println(num1)
	//2 Go程间channel访问顺序问题造成死锁
	//使用channel一端读或写，要保证另一端同时具有写或读的机会
	c2 := make(chan int)
	fmt.Println(<-c2) // <- c2 读不了，因为里面没有内容，会阻塞
	go func() {
		c2 <- 123
	}()
	//3 多Go程，交叉死锁
	// 自己控制一个条件，却还想要另一个条件
	c3 := make(chan int)
	c4 := make(chan int)
	go func() {
		for {
			select {
			case num1 := <-c3:
				{
					c4 <- num1
				}
			}
		}
	}()
	for {
		select {
		case num1 := <-c4:
			{
				c3 <- num1
			}
		}
	}

	//4. 在Go程中，尽量不要把互斥锁和读写锁，与channel混用，容易造成隐形死锁
}
