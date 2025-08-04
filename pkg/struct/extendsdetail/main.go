package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A name is", a.Name)
}

func (a *A) hello() {
	fmt.Println("a say hello")
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("b name is", b.Name)
}

func main() {
	var b B
	b.A.Name = "tom"
	b.A.age = 19
	b.A.SayOk() // A say ok
	b.A.hello()

	b.Name = "smith"
	b.age = 20
	b.SayOk()   // hello B smith
	b.hello()   // smith
	b.A.SayOk() // tom
}
