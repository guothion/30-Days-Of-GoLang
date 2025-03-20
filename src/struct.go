package main

import "fmt"

func main() {}

func init() {
	type newInt int
	type myInt = int // 别名

	var a newInt = 1
	var b myInt = 2
	fmt.Printf("type of a: %T\n", a) // newInt
	fmt.Printf("type of b: %T\n", b) //  int
}

type person struct {
	name string
	city string
	age  int8
}

func init() {
	var p person
	p.name = "liming"
	p.city = "beijing"
	p.age = 10
	fmt.Printf("type of p: %T\n", p)
	fmt.Printf("p value is %v\n", p)
	fmt.Printf("value format of p: %#v\n", p)

	var p2 = new(person)
	p2.name = "chenduxiu"
	p2.city = "shanghai"
	p2.age = 22
	fmt.Printf("type of p2: %T\n", p2)
	fmt.Printf("p2 value is %v\n", p2)
	fmt.Printf("value format of p2: %#v\n", p2)

	p3 := &person{}
	fmt.Printf("type of p3: %T\n", p3)
	fmt.Printf("p3 value is %#v\n", p3)
	p3.name = "lirongrong"
	p3.city = "xian"
	p3.age = 44
	fmt.Printf("value format of p3: %#v\n", p3)
}
