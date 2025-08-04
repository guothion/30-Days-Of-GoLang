package main

import "fmt"

type Monkey struct {
	Name string
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树")
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

type LittleMonkey struct {
	Monkey
}

// 我们通过接口来随意扩展我们的结构体的功能，并且我们没有影响到定义的猴子类型
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习会飞行")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "通过学习会游泳")
}

func main() {
	// 创建一个 LittleMonkey 实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
