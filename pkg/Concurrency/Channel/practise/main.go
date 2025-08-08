package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启 goroutine 将 0～100的数发送到 ch1 中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		// 关闭 ch1
		close(ch1)
	}()

	// 开启goroutine 从ch1中接收值，并将该值的平方送到ch2中
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}
	// 从上面的例子中我们看到有两种方式在接收值的时候判断通道是否被关闭，我们通常使用的是for range的方式。
}
