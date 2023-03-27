package main

import (
	"fmt"
	"net"
)

//每个人的基本信息
type person struct {
	msgChan chan string
	name    string
	addr    string
}

//全部人员列表
var allActivePerson map[string]person

//全局消息
var allMessage chan string

func main() {

	//创建tcp
	listen, err := createTCPBaseSocket()
	if err != nil {
		return
	}

	// 创建管理者go程，管理map 和全局channel
	go managerAllChannel()

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
func managerAllChannel() {
	allActivePerson = make(map[string]person)

	// 监听全局channel 中是否有数据, 有数据存储至 msg， 无数据阻塞。
	for {
		msg := <-allMessage
		//把消息发给所有人的信道，让每个人自己的conn自行处理数据
		for _, v := range allActivePerson {
			v.msgChan <- msg
		}
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
	//1 保存新用户信息到全体成员列表
	addr := conn.RemoteAddr().String()
	p := person{
		msgChan: make(chan string),
		name:    addr,
		addr:    addr,
	}
	allActivePerson[addr] = p
	//2 创建一个Go程，专门用于发消息
	go func() {
		fmt.Println("go func")
		for {
			msg := <-allActivePerson[addr].msgChan // 阻塞自己的消息通道，用于自行处理数据

			conn.Write([]byte(msg + "\n"))
		}
	}()

	//2 通知所有人该用户上线了，写入这个用户的的信道里面
	allMessage <- p.name + " > 用户上线了"

}
