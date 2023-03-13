package main

import (
	"fmt"
	"time"
)

func main() {
	// func (对象 结构体类型) 方法名（参数列表）（返回值列表） ｛
	// 代码体
	// ｝
	log := Log{}
	log.PrintNowTime()
	timeLog := TimeLog{}
	(&timeLog).PrintNowTime() //
	timeLog.PrintNowTime()    // 简化写法，把&timeLog简化为timeLog
}

type Log struct {
}

func (log Log) PrintNowTime() {
	fmt.Println(time.Now())
}

type TimeLog struct {
}

func (timeLog *TimeLog) PrintNowTime() {
	fmt.Println(time.Now())
}
