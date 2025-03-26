前面的例子，只是最简单的格式化，使用`Printf`是完全足够的。但是有时候会需要复杂的打印格式，
这时候一般需要将格式化代码分离出来以便更安全地修改。这些功能是由`text/template`和`html/template`等
模板包提供的，它们提供了一个将变量值填充到一个文本或**HTML**格式的模板的机制。

一个模板是一个字符串或一个文件，里面包含了一个或多个由双花括号包含的`{{action}}`对象。
大部分的字符串只是按字面值打印，但是对于`actions`部分将触发其它的行为。每个`actions`都
包含了一个用模板语言书写的表达式，一个`action`虽然简短但是可以输出复杂的打印值，模板语言包含
通过选择结构体的成员、调用函数或方法、表达式控制流`if-else`语句和`range`循环语句，还有其
它实例化模板等诸多特性。下面是一个简单的模板字符串：

```go
package main
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`
```
生成模板的输出需要两个处理步骤。第一步是要分析模板并转为内部表示，然后基于指定的输入执行模板。
分析模板部分一般只需要执行一次。下面的代码创建并分析上面定义的模板`templ`。注意方法调用链的
顺序：`template.New`先创建并返回一个模板；`Funcs`方法将`daysAgo`等自定义函数注册到模板中，
并返回模板；最后调用`Parse`函数分析模板。

```go
package main

import (
	"html/template"
	"log"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
}
```
