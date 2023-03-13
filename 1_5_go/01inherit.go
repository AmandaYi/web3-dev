package main

import "fmt"

type BaseModule struct {
	id       int
	shopName string
	Price    int
}

type OrderModule struct {
	BaseModule
	orderId string
}
type PayModule struct {
	BaseModule BaseModule
	payId      string
}

type PayModuleDescription struct {
	descriptionId int
	PayModule
}
type PayModuleDescriptionLog struct {
	PayModuleDescription
}

func main() {
	var order OrderModule = OrderModule{BaseModule: BaseModule{id: 100, shopName: "巧克力", Price: 10}, orderId: "200"}

	var payOrder PayModule = PayModule{BaseModule: BaseModule{id: 100, shopName: "巧克力", Price: 10}, payId: "200"}

	fmt.Println(order)
	fmt.Println(payOrder)

	var payModuleDescription PayModuleDescription = PayModuleDescription{PayModule: PayModule{
		BaseModule: BaseModule{id: 200, shopName: "巧克力", Price: 20},
		payId:      "300",
	}}
	fmt.Println(payModuleDescription)
	fmt.Println(payModuleDescription.BaseModule.shopName, payModuleDescription.BaseModule.id, payModuleDescription.BaseModule.Price)

	var payModuleDescriptionLog PayModuleDescriptionLog = PayModuleDescriptionLog{PayModuleDescription{
		descriptionId: 1000,
		PayModule: PayModule{BaseModule: BaseModule{
			id:       300,
			shopName: "巧克力",
			Price:    30,
		}, payId: "200"},
	}}
	fmt.Println(payModuleDescriptionLog.PayModuleDescription.BaseModule.id)
}
