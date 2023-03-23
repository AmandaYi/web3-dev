package main

import (
	"fmt"
	"time"
)

func main() {
	//select用来监听channel的数据流通，每一个case均为一个IO操作

	var userLogin chan bool = make(chan bool)
	go func() {
		time.AfterFunc(time.Second, func() {

			fmt.Println("登录成功!")

			defer func() {
				userLogin <- true
			}()
		})
	}()

	//这里的AfterFunc内部其实创建了一个Go程，因此不会进行阻塞代码
	time.AfterFunc(time.Second*5, func() {
		fmt.Println("5秒后用户再次登录了")
		userLogin <- true
	})

	var userLogout chan bool = make(chan bool)
	//这里的AfterFunc内部其实创建了一个Go程，因此不会进行阻塞代码
	time.AfterFunc(time.Second*10, func() {
		fmt.Println("6秒后用户退出了")
		userLogout <- true
	})

	defer func() {
		fmt.Println("系统已关机")
	}()
	var selectBreak bool
	for {
		time.Sleep(time.Second)

		if selectBreak == true {
			break
		}
		select {
		case result := <-userLogin: // 如果没有default，那么这里会一直阻塞
			{

				fmt.Println("用户登录了记录一下日志", result) // 用户登录了记录一下日志 true
			}
		case <-userLogout:
			{
				fmt.Println("发现用户推出了，系统也退出")
				selectBreak = true
				break
			}
			// select超时处理
		case <-time.After(time.Second * 10): // 只要userLogout的执行秒数大于这里的10秒，也就是说明了userLogout在很长时间后才触发，因此这里就会进入超时处理
			{

				fmt.Println("已超过10秒响应时间，正在退出程序")
				selectBreak = true
				break
			}

			//default: // 如果有default，那么select就不会阻塞，会直接把default内部的内容执行后，接着跳出select，往下执行代码
			//	//一般不写default，否则容易造成忙轮询，消耗资源
			//	{
			//		fmt.Println("用户系统运行中")
			//	}
		}
	}

}
