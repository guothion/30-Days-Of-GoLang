package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func(s string) {
		for i := 0; i < 3; i++ {
			fmt.Println(s, i)
		}
	}("world")

	// 主协程
	for i := 0; i < 3; i++ {
		// 切一下，再次分配任务
		// runtime.Gosched会让出 CPU ，让其他可运行的 goroutine 执行
		runtime.Gosched()
		fmt.Println("hello")
	}
}
