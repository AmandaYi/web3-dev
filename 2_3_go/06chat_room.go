package main

import (
	"fmt"
	"net"
)

//全部人员列表
var allPerson map[string]person

//每个人的基本信息
type person struct {
	name string
	addr string
}

func main() {
	//创建tcp
	listen, err := createTCPBaseSocket()
	if err != nil {
		return
	}
	for {
		//监听连接，三次握手
		conn, err := createConnAcceptSocket(listen)
		if err != nil {
			return
		}
		//处理用户逻辑
		chatHandle(conn)
	}
}

func createTCPBaseSocket() (net.Listener, error) {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("基本Socket创建失败", err)
		return nil, err
	}
	return listen, nil
}
func createConnAcceptSocket(listener net.Listener) (net.Conn, error) {
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("客户端建立连接失败", err)
		return nil, err
	}
	fmt.Println("客户端建立连接成功")
	return conn, nil
}
func chatHandle(conn net.Conn) {
	//打印用户登录成功信息
	fmt.Println("用户登录成功", conn.RemoteAddr().String())
	//通知所有的人

}
