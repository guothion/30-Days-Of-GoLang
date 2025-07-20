**map**是一种无序的基于`key-value`的数据结构，Go语言中的**map**是**引用类型**，**必须初始化才能使用**。
- **map**中所有的`key`都有相同的类型，`key`必须是支持`==`比较运算符的数据类型
- 所有的`value`也有着相同的类型
- 但是`key`和`value`之间可以是不同的数据类型。
- **map**的默认值也是`nil`
> map[**KeyType**]**ValueType**

## 创建
通过`make`函数来创建**map**，创建出来的是一个空的**map**。
> make(map[KeyType]ValueType,[cap])

其中`cap`表示**map**的容量，该参数虽然不是必须的，但是我们应该在初始化**map**的时候就为其指定一个合适的容量。
```go
package main
var ages = make(map[string]int)
```
我们也可以用**map字面值**的语法创建**map**,并初始化：
```go
package main
var ages = map[string]int{
	"alice": 31,
	"charlie": 40,
}
```
相当于，我们`make`创建一个空的**map**，并且给他赋值：
```go
package main
var ages = make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 40
```
## 判断某个值是否存在
Go语言中有个判断**map**中键是否存在的特殊写法，格式如下:
> value,ok := map[key]

```go
package main

import "fmt"

func init() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
        fmt.Println("查无此人")
	}

	if v, ok := scoreMap["bob"]; !ok { /* ... */ }
}
```
## 删除某个键值
我们使用`delete`来删除**map**的键值。
> delete(map,key)

```go
package main
var ages = map[string]int{
	"alice": 31,
	"charlie": 40,
}
delete(ages,"alice")
```
## 禁止对map中的元素指针操作
**map**中的元素并不是一个变量，因此我们不能对**map**的元素进行取址操作。

禁止对**map**元素取址的原因是**map**可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。

```go
package main
var ages = map[string]int{
	"alice": 31,
	"charlie": 40,
}
_ = &ages["bob"] // compile error: cannot take address of map element
```

## 遍历map(重点)
使用`range`来实现遍历。
```go
package main
func init() {
	ages := map[string]int{
		"alice": 31,
		"charlie": 40,
    }
//	遍历语法
    for key,val := range ages {
		fmt.Println(key, val)
    }
}
```
**注意**： 遍历**map**时的元素顺序与添加键值对的顺序无关。
> **Map**的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。在实践中， 遍历的顺序是随机的，
> 每一次遍历的顺序都不相同。这是故意的，每次都使用随机的遍历顺序可以 强制要求程序不会依赖具体的哈希函数
> 实现。如果要按顺序遍历**key/value**对，我们必须显式 地对**key**进行排序，可以使用**sort**包的**Strings**函数对字
> 符串**slice**进行排序。

下面是我们常用的方式：
```go
package main

import (
	"fmt"
	"sort"
)

//var names = []string

func main() {
	// 定义 map
	ages := map[string]int{
		"alice":   31,
		"charlie": 40,
	}
	// 我们可以创建一个容量和 map 一样的
	names := make([]string,0,len(ages))

	for _, name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n",name,ages[name])
	}
}
```

## 实践

### 实现一个 Set 
我们都知道Golang中没有 Set 类型，但是我们可以通过 **map** 实现一个类似的功能。

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
```
后边我们会讲**位向量**，我们通过map的方式来实现**Set**功能有时候不是最佳的，对于不同的情况，有不同的处理办法。
### 元素为map类型的切片

```go
package main

import "fmt"

func main() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d,value:%v",index,value)
	}
	fmt.Println("after init")
	// 对切片中的元素初始化
	mapSlice[0] = make(map[string]string,10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index,value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
    }
}
```
### 值为切片类型的map
```go
package main

func main() {
	sliceMap := make(map[string][]string,3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
```
### 对于不可以比较的类型作为 key
有些时候我们需要一个map，并且需求它的键是 slice ，但是因为 map 的键必须是可以比较的，所以直接实现肯定是不行。
我们可以通过一个函数，来映射我们的 slice 到 字符串，并且对于相同的 slice 最终得到的结果也是相同的。
```go
var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}
 func Add(list []string) {
	 m[k(list)]++
}
func Count(list []string) int {
	return m[k(list)]
}
```
> 这里我们需要学习 `%q` 这种写法。格式化动词（format verb）。它的作用是对值进行 引号转义（quoted） 输出，使其更易读，特别是对特殊字符。
> 转义字符在这个格式里边不会起到作用，并且引号仍然会存在。

当然上边只是一个示例，总之思想就是做一个函数，将不能比较的类型转化为可以比较的字符串，然后再放入map中。

