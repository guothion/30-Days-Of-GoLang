package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Name = "小明"
	p1.Age = 10
	var p2 *Person = &p1

	fmt.Println((*p2).Name)
	fmt.Println(p2.Name)
	p2.Name = "Tom~"
	fmt.Printf("p2.Name=%v p1.Name=%v \n", p2.Name, p1.Name)
	fmt.Printf("p2.Name=%v p1.Name=%v \n", (*p2).Name, p1.Name)

	fmt.Printf("p1的地址:%p\n", &p1)
	fmt.Printf("p2的地址:%p\np2的值是 %p", &p2, p2)
}
