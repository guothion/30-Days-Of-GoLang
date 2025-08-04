package main

import "fmt"

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

// 编写一个方法改变 i 的值
func (i *integer) change() {
	*i = *i + 1
}

func main() {
	var i integer = 10
	i.print()
	i.change()
	i.print()
}
