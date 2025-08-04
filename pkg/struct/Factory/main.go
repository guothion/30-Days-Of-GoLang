package main

import (
	"fmt"
	"src/pkg/struct/Factory/model"
)

func main() {
	// 创建一个 Student 的实例
	var stu = model.NewStudent("tom~", 88)
	fmt.Println(stu, *stu)
	fmt.Println(stu.Name, stu.GetScore())
}
