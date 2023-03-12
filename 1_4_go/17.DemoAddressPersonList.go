package main

import "fmt"

type Person struct {
	name    string
	contact []map[string]string
}

var PersionList []Person = []Person{Person{name: "zzy", contact: []map[string]string{map[string]string{"个人": "123456"}}}}

func main() {

}

func init() {
	for {
		showMain()
	}
}

func showMain() {
	fmt.Println("1：添加联系人")
	fmt.Println("2：删除联系人")
	fmt.Println("3：查询联系人")
	fmt.Println("4：编辑联系人")

	var no int
	_, err := fmt.Scanf("%d", &no)
	if err != nil {
		fmt.Println("请输入正确的选项重试")
	} else {
		DispatchType(no)
	}
}

func DispatchType(n int) {
	switch n {
	case 1:
		//添加联系人
		CreatePerson()
	case 2:
		//删除联系人
		DeletePerson()
	case 3:
		// 查询联系人
		FindPerson()
	case 4:
		//编辑联系人
		ModifyPerson()
	}
}

func CreatePerson() {
	var name string
	var contactKey string
	var contactValue string
	fmt.Println("请输入添加的姓名")
	fmt.Scan(&name)
	contact := make([]map[string]string, 0)
	for {
		fmt.Println("请选择添加的电话类型(学校，公司，个人)，添加完成请输入w")
		fmt.Scan(&contactKey)
		if contactKey == "w" {
			break
		}
		fmt.Printf("请输入%s号码\n", contactKey)
		fmt.Scan(&contactValue)
		contact = append(contact, map[string]string{contactKey: contactValue})
	}

	//保存起来
	PersionList = append(PersionList, Person{name: name, contact: contact})
	fmt.Println("当前电话列表如下：")
	showPersonList()
}
func DeletePerson() {
	fmt.Println("请输入您要删除的姓名")
	var deleteName string
	fmt.Scan(&deleteName)
	var deleteIndex = -1
	for i := 0; i < len(PersionList); i++ {
		if PersionList[i].name == deleteName {
			deleteIndex = i
		}
	}

	PrexList := PersionList[:deleteIndex]
	NextList := PersionList[deleteIndex+1:]
	PrexList = append(PrexList, NextList...)
	PersionList = PrexList
	fmt.Println("当前电话列表如下：")
	showPersonList()
}
func FindPerson() {
	fmt.Println("请输入您要查找的姓名")
	var findName string
	fmt.Scan(&findName)
	p := selectPerson(findName)

	if p != nil {
		fmt.Println("姓名：", p.name)
		fmt.Println("联系方式：")
		for _, mapItem := range p.contact {
			for phoneName, phone := range mapItem {
				fmt.Printf("电话（%s）是%s\n", phoneName, phone)
			}

		}
	} else {
		fmt.Printf("无法找到您输入的姓名%s\n", findName)
	}

}
func ModifyPerson() {
	//1. 分为编辑姓名和编辑电话
	fmt.Println("请输出您要编辑的姓名")
	var modifyName string
	fmt.Scan(&modifyName)
	p := selectPerson(modifyName)
	var contactKeyList []string
	if p != nil {

		for {
			fmt.Println("编辑姓名请输入1，编辑电话请输入2，返回请输入Q")
			var no string
			fmt.Scan(&no)
			switch no {
			case "1":
				fmt.Println("请输入新的姓名")
				fmt.Scan(&p.name)
			case "2":
				fmt.Println("当前的电话如下：")
				var i int = 0
				for _, mapItem := range p.contact {
					for phoneName, phone := range mapItem {
						fmt.Printf("电话（%s）是%s,编号：%d \n", phoneName, phone, i)
						contactKeyList = append(contactKeyList, phoneName)
						i++
					}
				}

				if len(contactKeyList) > 0 {

					fmt.Println("请选择您要修改的电话编号")
					var iScanIndex int
					fmt.Scan(&iScanIndex)

					var phone string
					fmt.Println("请输入要修改的新的号码")
					fmt.Scan(&phone)
					phoneName := contactKeyList[iScanIndex]

					p.contact[iScanIndex][phoneName] = phone
				}
			}

			if no == "Q" {
				break
			}
		}

	} else {
		fmt.Printf("无法找到您输入的姓名%s\n", modifyName)
	}
}

//辅助函数查询
func selectPerson(findName string) *Person {
	var p *Person
	for k, v := range PersionList {
		if v.name == findName {
			p = &PersionList[k]
		}
	}
	return p
}

// 展示出来所有的信息
func showPersonList() {
	if len(PersionList) == 0 {
		fmt.Println("列表无数据")
		return
	}
	for i := 0; i < len(PersionList); i++ {
		fmt.Println("姓名：", PersionList[i].name)
		fmt.Println("联系方式：")
		for _, mapItem := range PersionList[i].contact {
			for phoneName, phone := range mapItem {
				fmt.Printf("电话（%s）是%s\n", phoneName, phone)
			}

		}
	}
}
