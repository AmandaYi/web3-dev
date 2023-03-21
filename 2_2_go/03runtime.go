package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime")

	fmt.Println(runtime.Compiler) // 当前编译链
	fmt.Println(runtime.GOARCH)   // 当前运行平台
	fmt.Println(runtime.GOOS)     // 当前操作系统

	fmt.Println(runtime.MemProfileRate) // MemProfileRate控制在内存配置文件中记录和报告的内存分配比例

	runtime.GC() // GC运行垃圾收集并阻塞调用方，直到垃圾收集完成。它也可以阻止整个程序。

	// GOMAXPROCS设置可以同时执行的CPU的最大数量，并返回以前的设置。它默认为runtime.NumCPU的值。如果n〈1，则不改变当前设置。当调度程序改进时，此调用将消失。
	runtime.GOMAXPROCS(4) //  返回值是打印上一次的设置的核心数，如果没有设置，那么会返回程序自动获取的值

	fmt.Println(runtime.GOROOT()) // GOROOT路径
	runtime.Goexit()              // 退出协程
	runtime.Gosched()             //  Gosched让出处理器，允许其他goroutine运行。它不会挂起当前的goroutine，所以执行会自动恢复。

	runtime.LockOSThread() // 所有init函数都在启动线程上运行。从init函数调用LockOSThread将导致在该线程上调用main函数。https://pkg.go.dev/runtime@go1.20.2#LockOSThread

	runtime.NumCPU() // NumCPU返回当前进程可用的逻辑CPU数。

	runtime.NumCgoCall() // NumCgoCall返回当前进程进行的cgo调用的数量。

	runtime.NumGoroutine() // NumGoroutine返回当前存在的goroutine的个数。
}
