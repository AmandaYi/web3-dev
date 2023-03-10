package main

import "fmt"

func main() {
	//输入账号，密码，邮箱进行注册
	Register()
}

func Register() {
	var username, password, email string

	fmt.Println("请输入账号密码邮箱，格式为username:xxx password:xxx email:xxx")
	n, err := fmt.Scanf("username:%s password:%s email%s", &username, &password, &email)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(n)
		var isOK bool = Check(username, password, email)

		if isOK == true {
			SendEmail(username, email)
		} else {
			fmt.Printf("请检查输入的值")
			Register()
		}
	}
}

func Check(username, password, email string) bool {
	var isOK bool
	if username != "" && password != "" && email != "" {
		isOK = true
	}
	return isOK
}

func SendEmail(username, email string) {
	fmt.Printf("%s注册成功,邮箱地址%s", username, email)
}
