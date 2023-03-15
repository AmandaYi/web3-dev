package main

import "fmt"

func main() {

	var m MobileDisk
	var u UDrive
	Read(&m)
	Write(&m)

	Read(&u)
	Write(&u)
}

//用多态模拟实现，将移动硬盘或者U盘插入到电脑上进行读写数据

type ReadWriteEr interface {
	Read()
	Write()
}

func Read(RW ReadWriteEr) {
	RW.Read()
}
func Write(RW ReadWriteEr) {
	RW.Write()
}

type MobileDisk struct {
}
type UDrive struct {
}

func (m *MobileDisk) Read() {
	fmt.Println("我是移动硬盘，我在读取电脑数据")
}
func (m *MobileDisk) Write() {
	fmt.Println("我是移动硬盘，我在写电脑数据")
}
func (u *UDrive) Read() {
	fmt.Println("我是U盘，我在读取电脑数据")
}
func (u *UDrive) Write() {
	fmt.Println("我是U盘，我在写入电脑数据")
}
