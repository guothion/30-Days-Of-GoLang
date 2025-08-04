package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	fmt.Printf("r1.leftUp.x 的地址=%p\nr1.leftUp.y 的地址=%p\nr1.rightDown.x 的地址=%p\nr1.rightDown.y 的地址=%p\n", &r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	// 这里我们看下在内存里边是不是连续分布的
	// 通过我们大学学习的数据结构，我们知道是连续的
	// 连续分布的数据有什么特点

	// r2 有 2 个 *Point， 这两个 *Point 类型本身的地址是连续的，但是他们指向的地址不一定是连续的
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("r2.leftUp 的地址=%p\nr2.rightDown 的地址=%p", r2.leftUp, r2.rightDown)

}
