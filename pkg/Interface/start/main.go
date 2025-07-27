package main

import "fmt"

// 定义一个接口
type Usb interface {
	// 声明了 2 个没有实现的方法
	Start()
	Stop()
}

// Phone 实现
type Phone struct{}

// 让我们的结构体实现 Usb 的接口方法
func (p Phone) Start() {
	fmt.Println("Phone is starting")
}
func (p Phone) Stop() {
	fmt.Println("Phone is stopping")
}

// Camera 实现
type Camera struct{}

func (p Camera) Start() {
	fmt.Println("Camera is starting")
}
func (p Camera) Stop() {
	fmt.Println("Camera is stopping")
}

type Computer struct {
}

// 我们声明一个方法，接收一个接口类型的变量
// 只要实现了 Usb 接口（所谓实现 USb 接口，就是实现了 Usb 接口声明的所有方法）
func (c Computer) Working(usb Usb) {
	// 通过接口来调用方法
	usb.Start()
	usb.Stop()
}

func main() {
	// 测试
	// 先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	// 关键点
	computer.Working(phone)
	computer.Working(camera)
}
