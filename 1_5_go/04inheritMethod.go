package main

import "fmt"

type Common struct {
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
	fmt.Println(c.name)
}

func main() {
	var user *User = &User{CommonUser: CommonUser{id: 1, name: "admin", avatar: "http://admin.png"}, language: "zh-cn"}
	user.PrintCommonUser()

}
