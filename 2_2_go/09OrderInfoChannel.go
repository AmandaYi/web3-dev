package main

//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	go Factory(genChan)
//
//	go Trans(transport)
//
//	time.Sleep(time.Second * 10)
//}
//
//// 一个工厂生产货物，存储到仓库里面，
////物流运输，从仓库取出来
//type OrderInfo struct {
//	id int
//}
//
//var storeHouse chan *OrderInfo = make(chan *OrderInfo, 100)
//var genChan chan<- *OrderInfo = storeHouse
//var transport <-chan *OrderInfo = storeHouse
//
//// 工厂传入写入的通道，存到仓库
//var Factory = func(orderInfo chan<- *OrderInfo) {
//	i := 0
//	for {
//		i++
//		var genOrderInfo OrderInfo = OrderInfo{id: i}
//		orderInfo <- &genOrderInfo
//		time.Sleep(time.Second)
//	}
//}
//
//// 物流运输，取出仓库的货物
//var Trans = func(orderInfo <-chan *OrderInfo) {
//	for {
//		fmt.Printf("id = %d 的用户下单了", (<-orderInfo).id)
//	}
//}
