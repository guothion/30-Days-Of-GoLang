package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 方式一
	var p1 Person
	// 方式二
	p2 := Person{}

	fmt.Println(p1, p2)

	// 方式三
	var p3 *Person = new(Person)
	// 因为 p3 是一个指针，因此我们赋值方式得改下
	(*p3).Name = "smith" // 等价于 p3.Name = "smith"
	p3.Age = 30          // 等价于 (*p3).Age = 30
	// 这个位置完全是因为 golang 的设计者为了让我们使用方便
	fmt.Println(p3)
	p3.Name = "John Smith"
	fmt.Println(p3)

	// 方式四
	var p4 *Person = &Person{
		Name: "zhangdaqing",
		Age:  22,
	}
	fmt.Println(p4)
	p4.Name = "ximalaya" // 和上边一样，底层会对这里处理
	fmt.Println(p4)
}
