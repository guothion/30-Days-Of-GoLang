package main

import "fmt"

func main() {

	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	// 这里我们取切片 4-6项  4: "April", 5: "May", 6: "June"
	Q2 := months[4:7]
	// 这里我们取切片 6-8项  6: "June", 7: "July", 8: "August"
	summer := months[6:9]

	fmt.Println(Q2, len(Q2), cap(Q2))             // 4: "April", 5: "May", 6: "June"；长度 = 3； 容量 = 12 - (4-1) = 9
	fmt.Println(summer, len(summer), cap(summer)) // 6: "June", 7: "July", 8: "August"；长度=3；容量=12-5=7
	// fmt.Println(Q2[4], summer[4])
	endlessSummer := summer[:5] // extend a slice (within capacity)  6: "June", 7: "July", 8: "August",9: "September", 10: "October"
	fmt.Println(endlessSummer)  // "[June July August September October]"

	s := make([]string, 3) // len = cap = 3 的 内部是 string 类型的切片
	fmt.Println("emp:", s)
	fmt.Println(cap(s), len(s))

	d := make([]int, 5, 6)                    // int 类型切片 len = 5 cap = 6
	fmt.Println(cap(d), len(d), d)            // 6 5 [0 0 0 0 0]
	d = append(d, 1)                          // len = cap = 6
	fmt.Println(cap(d), len(d), d)            // 6 6  [0 0 0 0 0 1 ]
	fmt.Printf("%p\n里边的第一项指针%p\n", &d, &d[0]) // 我们查看下当前指针

	d = append(d, 2)                          // 此刻，容量超出，我们会额外创建一个新的切片来存放，和之前的不是一个指针地址
	fmt.Printf("%p\n里边的第一项指针%p\n", &d, &d[0]) // 我们查看扩容之后的指针
	// 总结：当容量超出之后，我们原本的切片的指针不会变，但是我们里边的内容，会重新指向最新的地址，这里应该是复制了一份原来的内容，存放到新的地址
	// 此刻，当我们修改切片内容项的时候不会影响到最初的内容，这点对于我们依赖某个数组创建切片导致扩容情况比较容易忽视。

	// fmt.Println(d)
	fmt.Println(cap(d), len(d), d) // 12 7  [0 0 0 0 0 1 2]

	e := []int{1, 2, 3, 4, 5}
	fmt.Println(cap(e), len(e), e) // 5 5 [1 2 3 4 5]

	f := make([]int, len(e))
	copy(f, e)
	fmt.Println("copy: ", f, cap(f)) // copy:  [1 2 3 4 5]

	l := d[2:5]
	fmt.Println(l, len(l), cap(l))

	//切片再切片
	a := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a:%v type:%T len:%d  cap:%d\n", a, a, len(a), cap(a))
	b := a[1:3]
	fmt.Printf("b:%v type:%T len:%d  cap:%d\n", b, b, len(b), cap(b))
	c := b[1:5]
	fmt.Printf("c:%v type:%T len:%d  cap:%d\n", c, c, len(c), cap(c))

}

func init() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}

func init() {
	source_arr := [7]string{
		"sunday",
		"monday",
		"tuesday",
		"wednesday",
		"thursday",
		"friday",
		"saturday",
	}
	work_day := source_arr[1:6]
	work_day[1] = "星期一"
	fmt.Println(work_day)
	fmt.Println(source_arr)

	work_day = append(work_day, "测试一", "测试二") // 此刻，已经超出了 cap
	fmt.Println(work_day)
	work_day[2] = "星期二"
	fmt.Println(work_day)
	fmt.Println(source_arr)
}
