package main

import "fmt"

type Base struct {
	id uint
}
type Common struct {
	Base
	activeUser bool
}
type CommonUser struct {
	id     uint
	name   string
	avatar string
	Common
}
type User struct {
	CommonUser
	language string
}

func (c *Common) SelectIsActiveUser() {
	fmt.Println(c.activeUser)
}
func (c *CommonUser) PrintCommonUser() {
	fmt.Println("CommonUser PrintCommonUser", c.name)
}
func (u *User) PrintCommonUser() {
	fmt.Println("User PrintCommonUser", u.name)
}
func main() {
	// var user *User = &User 和 var user User = User 是没有任何区别的，带（*）的是把User实例化后的地址存起来给user，不带（*）的是直接指向了用user指向了实例化的空间
	var user User = User{CommonUser: CommonUser{id: 1, name: "admin", avatar: "http://admin.png"}, language: "zh-cn"}

	// 如果有继承里面2个同名方法，那么会调用最近的方法
	user.PrintCommonUser()
	user.CommonUser.PrintCommonUser()
	// 方法值
	f1 := user.CommonUser.PrintCommonUser
	//f1 = (&user).CommonUser.PrintCommonUser // 可以用&user也可以不用，因为go会自动翻译
	fmt.Printf("f1的类型是%T\n", f1)

	//方法表达式
	f2 := (*User).PrintCommonUser // 这里要和定义的类型要一致，如 func(u *User) 与此处的(*User).PrintCommonUser
	f2(&user)                     // 这里要传入实例化后的对象，也要传入实例化对象的地址
	fmt.Printf("%T", f2)
}
