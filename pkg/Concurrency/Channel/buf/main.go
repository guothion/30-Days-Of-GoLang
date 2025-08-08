package main

import "fmt"

func recv(ch chan int) {
	ret := <-ch
	fmt.Println("接收成功", ret)
}

func main() {
	ch := make(chan int, 1)
	go recv(ch)
	ch <- 20
	fmt.Println("发送成功")
}
