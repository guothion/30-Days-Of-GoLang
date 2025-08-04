package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	// 定义一个结构体变量
	var p1 Person
	fmt.Println(p1) // { 0 [0 0 0 0 0] <nil> [] map[]}
	if p1.ptr == nil && p1.slice == nil && p1.map1 == nil {
		fmt.Println("p1.ptr is nil")
	}
	// 使用切片
	p1.slice = make([]int, 10)
	p1.slice[0] = 10
	// 使用map
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "value1"

	// 两个结构体互相不会影响
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1 // 值拷贝
	fmt.Println(monster2, monster1)
	monster2.Name = "青牛精"
	fmt.Println(monster2, monster1)
}
