# go 语言的主要特征
- 自动立即回收
- 丰富的内置类型
- 函数多返回值
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程
- 反射
- 语言交互性

# 25个关键字
```
break     default      func     interface   select
case      defer        go       map         struct
chan      else         goto     package     switch
const     fallthrough  if       range       type
continue  for          import   return      var
```
# 37个保留关键字
```
Constants:      true    false   iota    nil

Types:      int     int8    int16   int32   int64
            uint    uint8   uint16  uint32  uint64  uintptr
            float32 float64 complex128  complex64
            bool    byte    rune    string  error

Functions:  make    len     cap     new     append      copy      close     delete
            complex real    imag
            panic   recover
```
掌握上边这些关键字的用法基本就掌握了本门语言

# 可见性约定
- 声明在函数内部，是函数的本地值，类似于private
- 声明在函数外部，是对当前包可见（包内所有.go文件都可见）的全局值，类似于protect
- 声明在函数外部，并且首字母是大写，是所有包可见的全局值，类似于public





