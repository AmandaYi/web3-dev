package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//每个人的基本信息
type person struct {
	msgChan chan string
	name    string
	addr    string
}

//全部人员列表
var allActivePerson map[string]person

//全局消息,channel必须初始化，否则的话，信道不存在，自然永远堵塞
var allMessage chan string = make(chan string)

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
		go chatHandle(conn)
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
	defer conn.Close()
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
		for {
			msg := <-allActivePerson[addr].msgChan // 阻塞自己的消息通道，用于自行处理数据
			conn.Write([]byte(msg))
		}
	}()
	//2 通知所有人该用户上线了，写入这个用户的的信道里面
	allMessage <- p.name + "用户上线了"
	keepSend := make(chan bool)  // 是否保持在线
	waitClose := make(chan bool) // 等待关闭当前连接
	//处理输入输出
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				if n == 0 {
					fmt.Println(p.name, "退出")
					return
				}
			}
			//如果改名
			//一个汉字3个字节
			if len(buf) > 8 && string(buf[:6]) == "改名" {
				newName := strings.Split(string(buf[:n]), " ")[1]
				tmpBuf := []byte(newName)
				tmpBuf = tmpBuf[:len(tmpBuf)-2]
				newName = string(tmpBuf)
				if newName != "" {
					newName = strings.Replace("\r\n", "", newName, 1)
					p.name = newName
					p.name = newName
					allActivePerson[p.addr] = p
					continue
				}
			} else {
				// 加上用户+消息，发给全局消息发送中心，让其下发给每一个用户的发送消息的信道函数里面去
				allMessage <- p.name + ":" + string(buf[:n])
			}
			keepSend <- true
		}
	}()
	//监听是否需要退出用户
	go func() {
		for {
			select {
			case <-keepSend:
				{
				}
			case <-time.After(time.Second * 60):
				{
					allMessage <- p.name + "超时离开了"
					delete(allActivePerson, p.name)
					waitClose <- true
					return
				}
			}
		}
	}()
	<-waitClose
	return
}
