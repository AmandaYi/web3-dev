package main

import (
	"fmt"
	"time"
)

type OrderInfo struct {
	id int
}

// 实现一个厂房生产货物后，把货物放入到仓库后，再用物流车把仓库的东西运输出去
var storehouse chan *OrderInfo = make(chan *OrderInfo, 100) // 仓库缓冲区
var writeForklift chan<- *OrderInfo = storehouse            // 叉车把生产者东西存起来存到仓库
var readTransport <-chan *OrderInfo = storehouse            // 把仓库里面的货物拿出来

// 车间生成, 把东西放入仓库，需要传入只写通道
var lineProduction = func(orderInfo chan<- *OrderInfo) {
	var i int = 0
	for {
		i++
		orderInfo <- &OrderInfo{id: i}
		fmt.Printf("用户id = %d下单了，货物已放入仓库缓存\n", i)
		time.Sleep(time.Microsecond)
	}
}

// 货车拿货，把东西运输出去 应该传入只读通道
var vehicleTransportation = func(orderInfo <-chan *OrderInfo) {
	for {
		fmt.Printf("用户id = %d 的货物运输开始\n", (<-orderInfo).id)
	}
}

func main() {
	go lineProduction(writeForklift)
	go vehicleTransportation(readTransport)
	time.Sleep(time.Second * 10)
}
