package main

import "fmt"

type Person struct {
	Name string
}

// 一旦定义了接受者，那么这个方法只能被这个对象变量调用
func (p Person) test() {
	//如果我们加一个这个
	p.Name = "wangtianfeng"
	fmt.Println("test()", p.Name) // wangtianfeng
}

func (p *Person) test2() {
	p.Name = "minglou"
	fmt.Println("test2()", p.Name) // minglou
}

func main() {
	var p Person = Person{
		Name: "Tom",
	}

	p.test()
	fmt.Println(p.Name) // Tom
	// 这里我们说明我们调用方法的时候传入的是一个copy之后的形式参数

	// 如果我们使用指针呢？
	var p2 *Person = &Person{"mingjing"}
	p2.test2()
	fmt.Println(*p2) // minglou
}
