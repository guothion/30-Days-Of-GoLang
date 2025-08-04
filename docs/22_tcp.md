## TCP 编程
TCP/IP(Transmission Control Protocol/Internet Protocol) 即传输控制协议/网间协议，是一种面向连接（连接导向）的、
可靠的、基于字节流的传输层（Transport layer）通信协议，因为是面向连接的协议，数据像水流一样传输，会存在黏包问题。

一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。
因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次链接就创建一个goroutine去处理。

这里我们主要使用的是 `net` 包，我们看下官网给出的示例，我们看下如何使用。
- https://pkg.go.dev/net

```go
package net
// Conn 的方法
type Conn interface {
    Read(b []byte) (n int, err error)
    Write(b []byte) (n int, err error)
    Close() error
    LocalAddr() Addr
    RemoteAddr() Addr
    SetDeadline(t time.Time) error
    SetReadDeadline(t time.Time) error
    SetWriteDeadline(t time.Time) error
}
```

#### Server
TCP服务端程序的处理流程：
> 1.监听端口
>
> 2.接收客户端请求建立链接
>
> 3.创建goroutine处理链接。
```go
package main
import (
	"net"
)
func main() {
	ln,err := net.Listen("tcp", ":8080")
	if err != nil {
		//handle error
	}
	for {
		conn,err := ln.Accept() // 建立连接
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}	
}
// 处理链接
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {}
}
```

### Client
TCP 客户端通行流程如下：
> 1.建立与服务端的链接
>
> 2.进行数据收发
>
> 3.关闭链接
```go
package main
import (
	"bufio"
	"fmt"
	"net"
)
func main () {
	conn, err := net.Dial("tcp","golang.org:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')	
	fmt.Println(status)
}

```

## UDP编程
UDP协议（User Datagram Protocol）中文名称是用户数据报协议，是OSI（Open System Interconnection，开放式系统互联）参考模型中一种无连接的传输层协议，
不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，但是UDP协议的实时性比较好，通常用于视频直播相关领域。
### server
```go
package server
import (
    "net"
)
func main() {
	// 创建服务
    listen, err := net.ListenUDP("udp", &net.UDPAddr{
        IP:   net.IPv4(0, 0, 0, 0),
        Port: 30000,
    })
    if err != nil {}
	// 关闭服务
    defer listen.Close()
	// 开始接收数据
    for {
        var data [1024]byte
        n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
        if err != nil {}
        _, err = listen.WriteToUDP(data[:n], addr) // 发送数据
        if err != nil {}
    }
}
```
### client
```go
package client
import (
	"fmt"
	"net"
)
func main() {
	// 建立连接
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {}
	// 关闭连接
	defer socket.Close()
	// 发送数据
	sendData := []byte("Hello server")
	_, err = socket.Write(sendData)
	if err != nil {}
	// 读取返回的数据
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil{}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
```
## TCP粘包
我们先看下粘包的情况。
```go
// socket_stick/client/main.go
package client
import (
	"fmt"
	"net"
)
func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:30000")
    if err != nil {
        fmt.Println("dial failed, err", err)
        return
    }
    defer conn.Close()
    for i := 0; i < 20; i++ {
        msg := `Hello, Hello. How are you?`
        conn.Write([]byte(msg))
    }
}
```
最终结果是
```shell
收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?
收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?
收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?
收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?
收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?
```
为什么会出现这种情况？
主要原因就是tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发。
> 1.由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。
>
> 2.接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据。当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。

**解决方案**
出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行**封包**和**拆包**的操作。
> 封包：封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入”包尾”内容)。包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。

我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。
```go
package proto
import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) ([]byte,error) {
	// 读取消息的长度，转换成为int32类型（占 4 个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg,binary.LittleEndian, length)
	if err != nil {
		return nil, err
    }
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian,[]byte(message))
	if err != nil{
		return nil,err
    }
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte,_:=reader.Peek(4) // 读取前 4 字节的内容
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
    }
	if int32(reader.Buffered()) < length + 4 {
		return "",err
    }
	// 读取整整的消息数据
	if int32(reader.Buffered()) < length+4 {
		return "",err
    }
	
	pack := make([]byte,int(4+length))
	_,err = reader.Read(pack)
	if err !=  nil {
		return "",err
    }
	return string(pack[4:]),nil
}
```
服务端代码如下：

```go
package server

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:",err)
			return
		}
		fmt.Println("收到 client 发来的数据：",msg)
	}
}
func main() {
    listen, err := net.Listen("tcp","127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:",err)
		return
    }
	defer listen.Close()
	for {
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err:",err)
			continue
        }
		go process(conn)
    }
}
```
客户端代码如下：
```go
package client
import (
	"fmt"
	"net"
	"github.com/guothion/30-Days-Of-GoLang/pkg/proto"
)
func main() {
	conn,err := net.Dial("tcp","127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed,err:",err)
		return
    }
	defer conn.Close()
	for i:=0; i< 20; i++ {
		msg := `Hello, hello.How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed,err:",err)
			return
        }
		conn.Write(data)
    }
}
```