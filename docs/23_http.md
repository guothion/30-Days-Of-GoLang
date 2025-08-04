## http编程
### web工作流程
- Web服务器的工作原理可以简单地归纳为
  - 客户机通过TCP/IP协议建立到服务器的TCP连接
  - 客户端向服务器发送HTTP协议请求包，请求服务器里的资源文档
  - 服务器向客户机发送HTTP协议应答包，如果请求的资源包含有动态语言的内容，那么服务器会调用动态语言的解释引擎负责处理“动态内容”，并将处理得到的数据返回给客户端
  - 客户机与服务器断开。由客户端解释HTML文档，在客户端屏幕上渲染图形结果

### HTTP 服务端
```go
package server
import (
	"fmt"
	"net/http"
)

func main() {
	// 单独写回调函数
	http.HandleFunc("/go", myHandler)
	// addr 监听的地址
	// handler 回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr,"连接成功")
	// 请求方式
	fmt.Println("method:",r.Method)
	// go
	fmt.Println("url:",r.URL.Path)
	fmt.Println("header:",r.Header)
	fmt.Println("body:",r.Body)
	// 回复
	w.Write([]byte("www.5lmh.com"))
}
```
