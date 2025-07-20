package main

import "fmt"

func main() {
	// 1. scanln
	//testScanln()
	// 2. scanf
	testScanF()
}

func testScanln() {
	var (
		name   string
		age    byte
		sal    float32
		isPass bool
	)
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Printf("姓名：%s\n年龄： %d\n薪水： %f\n 是否通过考试 %t\n", name, age, sal, isPass)
}

func testScanF() {
	var (
		name   string
		age    byte
		sal    float32
		isPass bool
	)
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名：%s\n年龄： %d\n薪水： %f\n 是否通过考试 %t\n", name, age, sal, isPass)
}
