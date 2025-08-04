package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu *Stu) Say() {
	fmt.Println("stu say()", stu.Name)
}

type integer int

func (i integer) Say() {
	fmt.Println("integer say()", i)
}

type BInterface interface {
	Hello()
	Run()
}

// Monster 实现了 AInterface 和 BInterface
type Monster struct{}

func (m Monster) Hello() {
	fmt.Println("monster hello()")
}
func (m *Monster) Run() {
	fmt.Println("monster hello()")
}
func (m Monster) Say() {
	fmt.Println("monster say()")
}

func main() {
	var stu Stu = Stu{"lidazhao"}
	var a AInterface = &stu // 注意这个细节
	a.Say()
	var integer integer = 10
	var b AInterface = integer
	b.Say()

	var monster Monster
	var a2 AInterface = &monster
	var b2 BInterface = &monster
	a2.Say()
	b2.Hello()
}
