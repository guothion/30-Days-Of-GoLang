package main

import "fmt"

func main() {
	var ages map[string]int

	// 此刻 map 是nil
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)

	// 如果我们直接赋值会报错：panic: assignment to entry in nil map
	//ages["carol"] = 20

	// 我们需要初始化才可以
	ages = make(map[string]int)
	ages["carol"] = 20

	// 判断某个键值是否存在
	if _, isok := ages["Bob"]; !isok {
		ages["Bob"] = 30
	}

	list1 := []string{"beijing", "shanghai"}
	list2 := []string{"xian", "wuhan", "xining"}
	list3 := []string{"changsha", "chongqing", "gaungzhou", "zhejiang"}
	Add(list1)
	Add(list2)
	Add(list2)
	Add(list3)
	Add(list3)
	Add(list3)
	fmt.Println(Count(list1))
	fmt.Println(Count(list2))
	fmt.Println(Count(list3))
}

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}
func Add(list []string) {
	m[k(list)]++
}
func Count(list []string) int {
	return m[k(list)]
}
