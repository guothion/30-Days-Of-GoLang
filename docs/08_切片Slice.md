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

## cap 底层分配原理
这里我们说一个结论，就是当我们的容量超了之后，cap会自动在当前容量的基础上乘 2。
- 如果切片的容量小于1024，那么新的容量将是原容量的两倍。
- 如果切片的容量大于等于1024，那么新的容量将以原容量的1.25倍递增。

## 切片拷贝
我们使用`copy(slice1,slice2)`来实现拷贝。需要注意的是，函数`copy`在两个`slice`之间复制数据，
复制长度以`len`小的为准。
```go
package main
import "fmt"

func main() {
	s1 := []int{1,2,3,4,5}
	fmt.Printf("slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3 : %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3 : %v\n", s3)
}
```
## 字符串和切片
按照我们之前讲的，如果我们要修改字符串需要先将字符串转为数组。字符串的底层其实就是一个字节数组，
如果是中文的就是`rune`数组，所以我们也可以执行切片操作。