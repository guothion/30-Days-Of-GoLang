package main

import (
	"fmt"
	"sync"
)

func recv(ch chan int) {
	ret := <-ch
	fmt.Println("接收成功", ret)
}

func recv1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ret := <-ch
	fmt.Println("接收成功", ret)
}

func main() {
	ch := make(chan int)

	// 情况一：
	//go recv(ch) // 启用 goroutine 从通道接收值
	//ch <- 10
	//fmt.Println("发送成功")

	// 情况二：
	//go recv(ch) // 启用 goroutine 从通道接收值
	//fmt.Println("还未发送")
	//ch <- 10
	//fmt.Println("发送成功")

	// 情况三：
	//go recv(ch)
	//fmt.Println("还未发送")
	//ch <- 10
	//fmt.Println("发送成功")
	//time.Sleep(time.Second)

	// 情况四：
	var wg sync.WaitGroup
	wg.Add(1)
	go recv1(ch, &wg)
	fmt.Println("还未发送")
	ch <- 10
	fmt.Println("发送成功")
	wg.Wait() // 等待 recv goroutine 执行完成

}
