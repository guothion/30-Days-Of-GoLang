package main

import "fmt"

// 如果想在两个结构体之间转换，两个结构体必须拥有相同的字段和类型

type A struct {
	Num int
}

type B struct {
	Num int
}

type C A // 这里我们定义了结构类型 C，虽然看似和 A 一样，但是是两种数据类型，如果我们需要赋值的话需要强制转换

func main() {
	var a A
	var b B
	fmt.Println(a, b)
	a = A(b) // 强制转换可以，但是结构体必须完全一样，名字，结构，类型，个数
	fmt.Println(a)

	var c C
	c = C(a) // 如果我们想把 a 赋值给 c 必须使用强制转换
	fmt.Println(c)

}
