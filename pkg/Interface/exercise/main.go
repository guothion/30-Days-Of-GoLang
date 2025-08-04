package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 我们希望也可以调用 sort 包提供的方法
/***
官网说是提供了这三个方法就可以使用了
	type Interface interface {
		// Len方法返回集合中的元素个数
		Len() int
		// Less方法报告索引i的元素是否比索引j的元素小
		Less(i, j int) bool
		// Swap方法交换索引i和j的两个元素
		Swap(i, j int)
	}
**/

type Hero struct {
	Name string
	Age  int
}

// 声明一个切片类型结构体
type HeroSlice []Hero

// 实现官网的三个方法
func (h HeroSlice) Len() int {
	return len(h)
}

// 按照年龄从小到大排序
func (h HeroSlice) Less(i, j int) bool {
	return h[i].Age < h[j].Age
}
func (h HeroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	//
	var intSlice = []int{0, -1, 10, 7, 90}
	// 要求对intSlice进行排序
	// 1.冒泡排序
	// 2.也可以使用系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 请大家对结构体进行排序
	var heroes HeroSlice
	for i := 0; i < 20; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄～%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	fmt.Println(heroes)
	sort.Sort(heroes)
	fmt.Println(heroes)
}
