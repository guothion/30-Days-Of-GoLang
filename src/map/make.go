package main

import (
	"fmt"
)

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func init() {
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}

func init() {
	var arr0 = [5]int{1, 2, 3}
	arr1 := [...]int{1, 2, 3, 4, 5, 6}
	arr2 := [5]string{1: "hello", 4: "world"}
	arr3 := [...]struct {
		name string
		age  int
	}{
		{"liming", 12},
		{"wangli", 22},
	}
	fmt.Println(arr0, arr1, arr2, arr3)
}

func main() {
	ages := make(map[string]int, 20)
	fmt.Println(ages)
	ages["liming"] = 20
	ages["wanghong"] = 30
	fmt.Println(ages)
	for kv, ok := range ages {
		fmt.Println(kv, ok)
	}
}

func myTest(a [5]int, target int) {
	// 遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func init() {
	b := [5]int{1, 3, 5, 8, 7}
	myTest(b, 8)
}
