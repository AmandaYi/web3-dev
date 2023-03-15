package main

import "fmt"

func main() {
	// 接口就是一种规范的标准，只是规定了要做哪些事情，具体怎么做，接口是不管的，接口把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口

	var webUser WebUser
	var userCenterer UserCenterer
	userCenterer = &webUser
	userCenterer.Login()

	var adminUser AdminUser
	userCenterer = &adminUser
	userCenterer.Login()

	//结构体都实现了此接口，初始化的结构体把自己原始结构体地址给了接口，然后接口调用一样的方法，却有不同的逻辑，这种就叫做多态
	//所谓多态，指的是多种表现形式，多态就是同一个接口，使用不同的实例而执行不同操作
}

//定义接口的时候，最好以er结尾
type UserCenterer interface {
	Login()
	Register()
}

type WebUser struct {
}
type AdminUser struct {
}

func (webUser *WebUser) Login() {
	fmt.Println("前端web登录")
}
func (webUser *WebUser) Register() {
	fmt.Println("前端web注册")
}
func (adminUser *AdminUser) Login() {
	fmt.Println("admin web登录")
}
func (adminUser *AdminUser) Register() {
	fmt.Println("admin web注册")
}
