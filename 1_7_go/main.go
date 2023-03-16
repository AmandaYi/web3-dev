package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

type Window interface {
	ShowWindow() // 展示窗口
}
type ComWindow struct {
	Window
	*walk.MainWindow
}

func (comWindow *ComWindow) ShowWindow() {
	pathWindow := new(ComWindow)
	mainWindow := declarative.MainWindow{
		AssignTo: &pathWindow.MainWindow, // 关联主窗体
		Title:    "新窗口",
		MinSize:  declarative.Size{Width: 480, Height: 300},
	}
	err := mainWindow.Create()
	if err != nil {
		fmt.Println(err)
		return
	}
	pathWindow.SetX(600)
	pathWindow.SetY(300)
	pathWindow.Run()
}

type LabWindow struct {
	Window
}

func Show(WindowType string) {
	var win Window
	switch WindowType {
	case "main_window":
		win = &(ComWindow{})
	case "lab_window":
		win = &(LabWindow{})
	default:
		fmt.Println("参数错误")
	}
	win.ShowWindow()
}
func main() {
	Show("main_window")
}
