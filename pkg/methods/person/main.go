package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) speak() {
	fmt.Printf("%s is a good man", p.Name)
}

func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Printf("%s 计算的结果是： %d\n", p.Name, res)
}

func (p Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Printf("%s 计算 1~%d的累加结果是： %d\n", p.Name, n, res)
}

func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var p Person = Person{"tom~"}
	p.speak()
	p.jisuan()
	p.jisuan2(90)
	res := p.getSum(90, 70)
	fmt.Printf("最终结果: %d\n", res)
}
