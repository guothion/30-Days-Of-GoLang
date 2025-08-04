package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		// 1. 掌握函数 ReadRune 解码 UTF-8 编码，并且返回三个值
		r, n, err := in.ReadRune() // 返回 rune、nbytes、error
		// 如果我们读取到文件结尾 io.EOF ctrl+D
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
		}
		// 表示无效或无法表示的 Unicode 字符的替换字符
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters \n", invalid)
	}
}

// 我们可以给上边的函数做修改来统计字母数字和其他在 Unicode分类中的字符数量
// 使用方法是 unicode.IsLetter
