package main

import "fmt"

// 我们使用 type 定义接口
type Sayer interface {
	Say()
}

type Mover interface {
	move(string)
}

type Cat1 struct{}
type Dog1 struct{}

func (d Dog1) Say() {
	fmt.Println("汪汪汪")
}

func (c Cat1) Say() {
	fmt.Println("喵喵喵")
}

func (c Cat1) move(name string) {
	fmt.Println(name + "撒娇卖萌")
}

func (c *Dog1) move(name string) {
	fmt.Println(name + "旋转跳跃后空翻")
}

func main() {
	//var x Sayer
	//
	//a := Cat1{}
	//b := Dog1{}
	//
	//x = a
	//x.Say()
	//
	//x = b
	//x.Say()

	var x Mover
	var wangcai = Dog1{}
	x = wangcai
	x.move("旺财")

}
