package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000") //  这里创建了一个socket客户端管道
	defer conn.Close()
	if err != nil {
		fmt.Println("客户端Client建立连接失败", err)
		return
	}
	//不断的输入
	go func() {
		for {
			scanfBuf := make([]byte, 1024)
			n, err := os.Stdin.Read(scanfBuf)
			if err != nil {
				fmt.Println("os.Stdin.Read err", err)
				continue
			}

			_, err = conn.Write(scanfBuf[:n])
			if err != nil {
				fmt.Println("客户端发送数据失败", err)
				return
			}
		}
	}()

	//不断地回显
	for {
		var buf []byte = make([]byte, 4090)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端得到回应出错", err)
			return
		}
		fmt.Println((string)(buf[:n]))
	}
}
