结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。

Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。Go语言中通过结构体的内嵌再配合接口比面向对象
具有更高的扩展性和灵活性。

## 自定义类型
除了我们讲的string、整型、浮点型、布尔等数据类型，Golang 也支持我们自己定义类型。

### 类型定义
语法：
> type 自定义别名 类型

比如：

```go
package main

import "fmt"

type myInt int

var a myInt = 12

func init() {
	fmt.Printf("%T\n",a) // myInt
}
```
我们定义了一个类型叫`myInt`，它其实就是和`int`一模一样，只不过我们给他改了个名。但是我们如果是输出他的类型，就是我们定义的`myInt`
### 类型别名
顾名思义，就是换了个名字，但实际上还是原始的类型。当我们打印类型的时候，输出的类型和别名无关。
其实之前我们就接触过，比如我们字符串中的字符，我们用的是`byte`和`rune`这俩，这两个的类型其实就是：
```go
package main
type byte = int8
type rune = int32
```
我们来看下：

```go
package main

import "fmt"

type myInt = int
var b myInt
func init() {
	fmt.Printf("type is %T\n",b) // int
}
```
## 定义结构体类型
这里我们开始定义结构体类型，首先明白我们是定义新的类型，不是类型别名。
语法：
> type 类型名 struct { 结构体 }

比如：

```go
package main

import "time"

// 这里我们就定义了一个员工的类型，里边我们定义了他的属性（是不是很像数据表啊）
type Employee struct {
	ID                  int
	Name,Address        string
	Dob                 time.Time
	Position            string
	Salary,ManagerID    int
}

var dilbert Employee
```
访问结构体里边的值我们使用`.`操作符来访问。
比如：
```
dilbert.Salary -= 5000 // demoted, for writing too few lines of code
```
或者骚一点的操作，对成员取地址，然后通过地址访问：
```
position := &dilbert.Position
*position = "Senior " + *position // promoted, for outsourcing to Elbonia
```
或者更骚一点的操作：
```
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
# 相当于
(*employeeOfTheMonth).Position += " (proactive team player)"
```
> ⚠️：如果结构体成员名字是以**大写字母开头**的，那么该成员就是导出的；
> 这是Go语言导出规则决定的。一个结构体可能同时包含导出和未导出的成员。

## 指针类型结构体
一个命名为S的结构体类型将不能再包含**S**类型的成员：因为一个聚合的值不能包含它自身。
（该限制同样适用于数组。）

但是**S**类型的结构体可以包含`*S`指针类型的成员，这可以让我们创建递归的数据结构，比如链表和树结构等。

我们还可以通过使用`new`关键字对结构体进行实例化，得到的是**结构体的地址**。

```
var p2 = new(person)
fmt.Printf("%T\n", p2)     //*main.person
fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
```
从打印的结果中我们可以看出**p2**是一个结构体指针。

> 需要注意的是在Go语言中支持对结构体指针直接使用.来访问结构体的成员。

```go
var p2 = new(person)
p2.name = "测试"
p2.age = 18
p2.city = "北京"
fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"测试", city:"北京", age:18}
```
使用`&`对结构体进行取地址操作相当于对该结构体类型进行了一次`new`实例化操作。
```go
p3 := &person{}
fmt.Printf("%T\n", p3)     //*main.person
fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
p3.name = "博客"
p3.age = 30
p3.city = "成都"
fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}
```
`p3.name = "博客"`其实在底层是`(*p3).name = "博客"`，这是Go语言帮我们实现的语法糖。

在下面的代码中，我们使用一个二叉树来实现一个插入排序：
```go
package main
type tree struct {
	value       int
	left,right  *tree
}

func Sort(values []int) []int {
	var root *tree
	for _,v := range values {
		root = add(root, v)
    }
	return appendValues(values[:0],root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values,t.right)
    }
	return values
}

func add(t *tree, value int) *tree {
	if t==nil {
		t = new(tree)
		t.value = value
		return t
    }
	
	if value < t.value {
		t.left = add(t.left,value)
    } else {
		t.right = add(t.right,value)
    }
	return t
}
```
## 初始化
- 使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
- 也可以对结构体指针进行键值对初始化
- 初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值
  - 1.必须初始化结构体的所有字段。
  - 2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
  - 3.该方式不能和键值初始化方式混用。
```go
package main
func init() {
	// 方式一
	p5 := person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	// 方式二
	p6 := &person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	// 方式三
	p8 := &person{
		"pprof.cn",
		"北京",
		18,
	}	
}

```

## 构造函数
Go语言的结构体没有构造函数，我们可以自己实现。

```go
package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
func init() {
	p9 := newPerson("pprof.cn", "测试", 90)
	fmt.Printf("%#v\n", p9)
}
```

## 方法和接收者
Go语言中的方法（**Method**）是一种作用于特定类型变量的函数。
这种特定类型变量叫做接收者（**Receiver**）。接收者的概念就类似于其他语言中的`this`或者`self`。

```
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```
> 1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的**第一个小写字母**，而不是**self**、**this**之类的命名。
> 例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
> 
> 2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
> 
> 3.方法名、参数列表、返回参数：具体格式与函数定义相同。

举个例子：

```go
package main

import "fmt"

type Person struct {
	name string
	age  int8
}

func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n",p.name)
}
func main() {
	p1 := NewPerson("测试",25)
	p1.Dream()
}
```
## 指针类型的接收者
