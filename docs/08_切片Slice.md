## 切片定义
首先我们要知道切片不是数组，也不是数组指针。
- 切片不需要指定长度
- 切片是引用类型
它通过内部指针和相关属性引用数组片段，来实现的变长的方案

## 创建切片的方式
### 1.声明切片
```go
package main
func init() {
	var s1 []int
	s2 := []int
	s3 := []int{1,2,3} // 初始化赋值
}
```
### 2.make 函数
```go
package main
func init() {
	//这我们也要知道，第一个参数是类型
	//第二个参数是长度，第三个参数是容量
	s3 := make([]string, 0, 8)
}
```
### 3.从数组中切片
```go
package main
func init() {
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	s1 := arr[3:6] // 前包后不包
	s2 := arr[4:]
	s3 := arr[:5]
	s4 := arr[:]
	s5 := arr[::5]
}
```

## len(slice) 和 cap(slice)
- len(slice) - 求当前切片的长度
- cap(slice) - 求当前切片的容量

## append(slice, tar)函数
使用`append(slice,tar)`来给切片来追加元素。**返回的是最新的slice对象**。
```go
package main
import (
	"fmt"
)
func main() {
	s1 := make([]int, 0 ,5)
	fmt.Printf("%p\n", &s1)
	
	s2 := append(s1,1)
	fmt.Printf("%p\n", &s2)
	
	fmt.Println(s1,s2)
}
```
这里我们可以通过查看指针来判断是否是同一个地址。是否发生了复制。


