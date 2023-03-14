package main

func main() {
	// 接口就是一种规范的标砖，只是规定了要做哪些事情，具体怎么做，接口是不管的，接口把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口

}

//定义接口的时候，最好以er结尾
type UserCenterer interface {
	Login()
	Register()
}
