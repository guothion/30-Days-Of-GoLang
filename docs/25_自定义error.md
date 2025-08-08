## 异常
### 异常处理
Golang 没有结构化异常，使用`panic`抛出错误，使用`recover`捕获错误。
> 使用场景：Go中可以抛出一个panic异常，然后在 defer 中通过 recover 捕获这个异常，然后正常处理。

> **panic**
> 1. 内置函数
> 2. 假如函数中写了`panic`语句，会终止代码。（如果`panic`所在函数中存在`defer`则，`defer`的正常执行）
> 3. 返回函数 F 的调用者G，在 G 中调用函数 F 语句之后的代码不会执行。
> 4. 直到`goroutine`整个退出，并报告错误

> **recover**
> 1. 内置函数
> 2. 用来控制一个`goroutine`的`panicking`行为，捕获`panic`，从而影响应用的行为
> 3. 一般的调用建议：
>    4. 在 `defer` 函数中，通过`recover`来终止一个`goroutine`的`panicking`过程，从而恢复正常代码执行
>    5. 可以获取通过`panic`传递的`error`

> **⚠️注意**
> 1. 利用`recover`处理`panic`指令，`defer`必须放在`panic`之前定义，另外`recover`只有在`defer`调用的函数中才有效。否则当`panic`时，`recover`无法捕获`panic`，无法防止`panic`扩散。
> 2. `recover`处理异常后，逻辑并不会恢复到`panic`那个点，函数跑到`defer`之后的那个点。
> 3. 多个`defer`会形成`defer`栈，后定义的`defer`语句会被最先调用。

-----------
向已经关闭的通道发送数据会引发`panic`
````go
package main
import (
	"fmt"
)
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
        }
    }()
	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- -1
}
````
------
延迟调用引发的错误，可被后续延迟调用捕获，但仅仅最后一个错误可以被捕获。（我们报错之后不影响`defer`调用，所以还会执行`defer`的报错）
```go
package main 
import "fmt"
func test() {
	defer func() {
		fmt.Println(recover())
    }()
	defer func() {
		panic("defer panic")
    }()
	
	panic("test panic")
}
func main() {
	test()
}
```
最终结果：`defer panic`。

-------
除了用`panic`引发中断性错误外，还可以返回`error`类型错误对象来表示函数调用状态。
```go
type error interface {
	Error() string
}
```
标准库`errors.New`和`fmt.Errorf`函数用于创建实现`error`接口的错误对象。通过判断错误对象实例来确定具体错误类型。
```go
package main

import (
	"errors"
	"fmt"
)

var ErrDivByZero = errors.New("division by zero")
func div(x,y int) (int,error) {
	if y== 0 {
		return 0,ErrDivByZero
    }
	return x/y,nil
}
func main() {
	defer func() {
		fmt.Println(recover())
    }()
	switch z,err := div(10, 0); err {
	case nil: println(z)
	case ErrDivByZero: panic(err)
	}
}
```
-------
我们实现一个 `try catch`
```go
package main
import "fmt"
func Try(fun func(), handler func(interface{})) {
	defer func(){
		if err := recover(); err != nil {
			handler(err)
        }
    }()
	fun()
}
func main() {
	Try(func() {
		panic("test panic")
    },func(err interface{}) {
		fmt.Println(err)
    })
}
```
### 自定义error
```go
package main
import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path string
	op string
	createTime string
	message string
}
// 核心
func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \n op=%s \ncreateTime=%s \nmessage=%s",p.path,p.op,p.createTime,p.message)
}

func Open(filename string) error {
	file,err := os.Open(filename)
	if err != nil {
		return &PathError{
			path: filename,
			op: "read",
			message: err.Error(),
			createTime: fmt.Sprintf("%v",time.Now()),
        }
    }
	
	defer file.Close()
	return nil
}

func main() {
	err := Open("/Users/5lmh/Desktop/go/test.txt")
	switch v:=err.(type) {
	case *PathError: fmt.Println("get path error", v)
	default:
	}
}
```