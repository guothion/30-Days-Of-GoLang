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

func init() {
	p4 := person{
		name: "wangsihui",
		city: "zhangjiakou",
		age:  25,
	}
	fmt.Printf("p4 value is %#v \n", p4)

	p5 := &person{
		name: "zhengwei",
	}
	fmt.Printf("p5 value is %#v\n", p5)

	//我们的键值其实是可以省略的，但是必须注意：
	//1.必须初始化结构体的所有字段。
	//2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	//3.该方式不能和键值初始化方式混用。
	p6 := &person{
		"guilong",
		"bejing",
		27,
	}
	fmt.Printf("p6 value is %#v\n", p6)
}

func init() {
	type student struct {
		name string
		age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	// 有问题的代码
	//for _, stu := range stus {
	//	m[stu.name] = &stu // 陷进：这里我们都指向了 stu 的指针，这个stu最后给的值是最后一个值
	//}

	// 修改后的代码
	for i, _ := range stus {
		m[stus[i].name] = &stus[i]
	}

	fmt.Printf("value of stu: %v\n", m)
	for k, v := range m {
		fmt.Printf("value of stu[%v]: %#v\n", k, v)
		fmt.Println(k, "=>", v.name)
	}

}

type Animal struct {
	name  string
	legs  int8
	color string
}

func newAnimal(name string, legs int8, color string) *Animal {
	return &Animal{name: name, legs: legs, color: color}
}

func (a *Animal) eat(food string) {
	fmt.Printf("%v eat %v", a.name, food)
}

func init() {
	woof := newAnimal("woof", 4, "grey")
	woof.eat("meat")
}
