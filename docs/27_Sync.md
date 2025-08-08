## Sync
### sync.WaitGroup
在代码中生硬的使用`time.Sleep`肯定是不合适的，Go语言中可以使用`sync.WaitGroup`来实现并发任务的同步。 

`sync.WaitGroup`有以下几个方法：

| 方法名                           | 功能              |
|-------------------------------|-----------------|
| `(wg *WaitGroup)Add(delta int)` | **计数器+Delta**   |
| `(wg *WaitGroup)Done()`         | **计数器-1**       |
| `(wg *WaitGroup)Wait()`         | **阻塞直到计数器变为 0** |

`sync.WaitGroup`内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了**N** 个
并发任务时，就将计数器值增加**N**。每个任务完成时通过调用`Done()`方法将计数器减**1**。
通过调用`Wait()`来等待并发任务执行完，当计数器值为**0**时，表示所有并发任务已经完成。

[示例](../pkg/Concurrency/Sync/waitGroup/example/main.go)

需要注意`sync.WaitGroup`是一个结构体，传递的时候要传递指针。

### sync.Once
在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。

Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–`sync.Once`。
`sync.Once`只有一个`Do`方法，其签名如下：
```go
func (o *Once) Do(f func()) {}
```
注意：如果要执行的函数`f`需要传递参数就需要搭配闭包来使用。

[示例](../pkg/Concurrency/Sync/Once/main.go)

### sync.Map
Go语言中内置的map不是并发安全的。
```go
var m = make(map[string]int)

func get(key string) int {
    return m[key]
}

func set(key string, value int) {
    m[key] = value
}

func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func(n int) {
            key := strconv.Itoa(n)
            set(key, n)
            fmt.Printf("k=:%v,v:=%v\n", key, get(key))
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```
上面的代码开启少量几个`goroutine`的时候可能没什么问题，当并发多了之后执行上面的代码就会报`fatal error: concurrent map writes`错误。

像这种场景下就需要为`map`加锁来保证并发的安全性了，Go语言的`sync`包中提供了一个开箱即用的并发安全版`map–sync.Map`。开箱即用表示不用像内
置的`map`一样使用`make`函数初始化就能直接使用。同时`sync.Map`内置了诸如`Store`、`Load`、`LoadOrStore`、`Delet`e、`Range`等操作方法。

```go
var m = sync.Map{}

func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func(n int) {
            key := strconv.Itoa(n)
            m.Store(key, n)
            value, _ := m.Load(key)
            fmt.Printf("k=:%v,v:=%v\n", key, value)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```