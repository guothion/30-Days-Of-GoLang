package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// 为了提高效率，我们更多的是使用指针
func (c *Circle) Area2() float64 {
	return math.Pi * (*c).radius * c.radius // 这里使用 (*c).redius 和 c.radius 都可以
}

type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name: %s, Age: %d", stu.Name, stu.Age)
	return str
}

func main() {
	cirle := Circle{10}
	fmt.Println(cirle.Area())

	cirle2 := Circle{5}
	fmt.Println((&cirle2).Area2()) // 标准的访问方式
	fmt.Println(cirle2.Area())     // 但是，我们这样写也可以，作者给我们处理了

	stu := Student{"James", 20}
	fmt.Println(&stu)
}
