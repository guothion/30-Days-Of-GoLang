package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("hello Goroutine", i)
}

// 这里我们看看并发执行的 goroutine 是不是按照顺序执行的
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait() // 等待所有的 goroutine 都结束
}

// 结论：并非按照顺序执行
