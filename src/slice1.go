package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	reverse(a[:])
	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5, 6}
	reverse(b[:2])
	reverse(b[2:])
	reverse(b)
	fmt.Println(b)
}

func reverse(s []int) {
	// 注意细节
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
