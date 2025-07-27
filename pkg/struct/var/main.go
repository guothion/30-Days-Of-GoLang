package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	// 在创建结构体变量的时候直接指定字段值
	stu1 := Student{"小明", 18}
	stu2 := Student{Age: 20, Name: "小雷"}
	// 下面是返回指针类型 ！！！
	stu3 := &Student{"王老师", 35}
	stu4 := &Student{
		Age:  55,
		Name: "黄老邪",
	}

	fmt.Println(stu1, stu2, stu3, stu4)
}
