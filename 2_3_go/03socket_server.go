package main

import (
	"bytes"
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:9000") // 这里也会有一个socket，用来绑定port+ip，
	// 作用1：用来复制一份新的socket用于通信，作用2，如果有新客户端来了，直接复制自身用于通信。
	if err != nil {
		fmt.Println("开启服务器错误", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept() // 这里创建了一个socket服务端管道
	if err != nil {
		fmt.Println("服务端监听任务启动失败", err)
		return
	}
	defer conn.Close()
	var buf []byte = make([]byte, 4090)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("服务端接受数据错误", err)
		return
	}
	fmt.Println("服务器接收到的数据是：", string(buf[:n]))
	_, err = conn.Write(bytes.ToUpper(buf[:n]))
	if err != nil {
		fmt.Println("服务端回应失败", err)
	}
}
