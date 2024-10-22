package main

import "fmt"

func main() {
	str := "hello 世界"
	for i, s := range str {
		fmt.Println(i, s)
	}

	fmt.Println("------------------------")

	j := 1
	for j < len(str) {
		s := str[j]
		fmt.Println(s)
		j++
		// defer j
	}
}
