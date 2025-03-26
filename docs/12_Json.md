JavaScript对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。
Go语言对于这些标准格式的编码和解码都有良好的支持，由标准库中的`encoding/json`、
`encoding/xml`、`encoding/asn1`等包提供支持。本节，我们将对重要的`encoding/json`包的用法做个概述。

### Marshal 方法
将一个Go语言中类似`movies`的结构体`slice`转为**JSON**的过程叫编组（**marshaling**）。
编组通过调用`json.Marshal`函数完成
```go
package main
import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}
func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
```
### MarshalIndent
为了生成便于阅读的格式，另一个`json.MarshalIndent`函数将产生整齐缩进的输出。
该函数有两个额外的字符串参数用于表示每一行输出的前缀和每一个层级的缩进：

```go
package main

import "fmt"

func init() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
```
### Unmarshal
编码的逆操作是解码，对应将**JSON**数据解码为Go语言的数据结构，Go语言中一般叫`unmarshaling`，
通过`json.Unmarshal`函数完成。下面的代码将**JSON**格式的电影数据解码为一个**结构体slice**，结构体中只有**Title**成员。

```go
package main

import "fmt"

func init() {
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
```