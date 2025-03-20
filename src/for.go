package main

import "fmt"

func main() {
	str := "hello 世界"
	for i, s := range str {
		fmt.Println(i, s) // s 是 rune 类型 Unicode
	}

	fmt.Println("------------------------")

	j := 1
	for j < len(str) {
		s := str[j] // s 是 byte 类型
		fmt.Println(s)
		j++
		// defer j
	}
}
