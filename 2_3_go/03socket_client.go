package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000") //  这里创建了一个socket客户端管道
	defer conn.Close()
	if err != nil {
		fmt.Println("客户端Client建立连接失败", err)
		return
	}
	s := "hello"
	_, err = conn.Write(([]byte)(s))
	if err != nil {
		fmt.Println("客户端发送数据失败", err)
		return
	}
	var buf []byte = make([]byte, 4090)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("客户端得到回应出错", err)
		return
	}
	fmt.Println((string)(buf[:n]))
}
