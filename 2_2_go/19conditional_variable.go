package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

//1. 创建 条件变量 var cond sync.Cond
var cCond sync.Cond

func main() {
	//条件变量

	//1 判断条件变量
	//2 枷锁
	//3 访问公共区
	//4 解锁
	//5 唤醒阻塞在条件变量上的对端

	//使用流程
	/*
		1. 创建 条件变量 var cond sync.Cond
		2. 指定条件变量用的锁: cond.L = new(sync.Mutex)
		3. cond.Lock() 给公共区域加锁
		4. 判断是否到达 阻塞条件（缓冲区满/空） --- for 循环判断
				for len(ch) == cap(ch) {  cond.Wait()  } // 1. 阻塞 2. 解锁 3.加锁
		5. 访问公共区 --- 读、写数据、打印
		6. 解锁条件变量用的锁 cond.Unlock()
		7. 唤醒阻塞在条件变量上的对端 cond.Signal()
	*/

	//2. 指定条件变量用的锁: cond.L = new(sync.Mutex)
	cCond.L = new(sync.Mutex)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 5; i++ {
		go CLineProduction(CWriteForklift)
	}
	for i := 0; i <= 5; i++ {
		go CVehicleTransportation(CReadTransport)
	}

	time.Sleep(time.Second * 10)
}

type COrderInfo struct {
	id       int
	shopName string
}

// 实现一个厂房生产货物后，把货物放入到仓库后，再用物流车把仓库的东西运输出去
var CStorehouse chan *COrderInfo = make(chan *COrderInfo, 100) // 仓库缓冲区
var CWriteForklift chan<- *COrderInfo = CStorehouse            // 叉车把生产者东西存起来存到仓库
var CReadTransport <-chan *COrderInfo = CStorehouse            // 把仓库里面的货物拿出来

// 车间生成, 把东西放入仓库，需要传入只写通道
var CLineProduction = func(orderInfo chan<- *COrderInfo) {
	for {
		id := rand.Intn(1000)

		//3. cond.Lock() 给公共区域加锁
		cCond.L.Lock()

		//4. 判断是否到达 阻塞条件（缓冲区满/空） --- for 循环判断
		/***************************重点***************************/
		//对于生产者而言，这里的条件变量一定是进行判断信道里面的值是不是满了，等于容量的大小的，才调用wait，进行等待生产，
		/***************************重点***************************/
		/***************************重点***************************/
		//这里使用for循环的原因是，当前的进行如果进行了wait后，一定会执行unlock释放，unlock同时，其他的Go程一定是在加锁的过程中进行等待，
		//当另一个突然拿到锁之后，会执行判断条件，然后进行调用，调用后发现已经满了，也会进行wait，从而释放锁，这样一来二去，让所有的生产者都进行了wait判断，从而阻塞全部的Go程
		/***************************重点***************************/
		/***************************重点***************************/
		//条件变量使用for循环的原因如上所示
		/***************************重点***************************/

		for len(orderInfo) == cap(orderInfo) {
			cCond.Wait() // 1. 阻塞 2. 解锁 3.加锁
		}

		//5. 访问公共区 --- 读、写数据、打印
		tmpOrderInfo := &COrderInfo{id: id, shopName: "No: " + strconv.Itoa(id) + "号产品"}
		orderInfo <- tmpOrderInfo
		fmt.Printf("用户id = %d下单了%s，货物已放入仓库缓存\n", tmpOrderInfo.id, tmpOrderInfo.shopName)

		//6. 解锁条件变量用的锁 cond.Unlock()
		cCond.L.Unlock()

		//7. 唤醒阻塞在条件变量上的对端 cond.Signal()
		cCond.Signal()

		time.Sleep(time.Microsecond)
	}
}

// 货车拿货，把东西运输出去 应该传入只读通道
var CVehicleTransportation = func(orderInfo <-chan *COrderInfo) {
	for {

		//3. cond.Lock() 给公共区域加锁
		cCond.L.Lock()

		//4. 判断是否到达 阻塞条件（缓冲区满/空） --- for 循环判断

		/***************************重点***************************/
		//对于消费者而言，这里的条件变量一定是进行判断信道里面的值是不是没有了，只有没有了才会进行调用wait阻塞，解锁，加锁
		/***************************重点***************************/
		for len(orderInfo) == 0 {

			cCond.Wait() // 1. 阻塞 2. 解锁 3.加锁
		}

		//5. 访问公共区 --- 读、写数据、打印
		tmpOrderInfo := <-orderInfo
		fmt.Printf("用户id = %d 的货物%s运输开始\n", tmpOrderInfo.id, tmpOrderInfo.shopName)

		//6. 解锁条件变量用的锁 cond.Unlock()
		cCond.L.Unlock()

		//7. 唤醒阻塞在条件变量上的对端 cond.Signal()
		cCond.Signal()
	}

}
