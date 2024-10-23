package main

import "fmt"

func main() {

	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2, len(Q2), cap(Q2))
	fmt.Println(summer, len(summer), cap(summer))
	// fmt.Println(Q2[4], summer[4])
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"

	s := make([]string, 3)
	fmt.Println("emp:", s)
	fmt.Println(cap(s), len(s))

	d := make([]int, 5, 6)
	fmt.Println(cap(d), len(d), d) // 6 5 [0 0 0 0 0]
	d = append(d, 1)
	d = append(d, 2)
	// fmt.Println(d)
	fmt.Println(cap(d), len(d), d) // 12 7  [0 0 0 0 0 1 2]

	e := []int{1, 2, 3, 4, 5}
	fmt.Println(cap(e), len(e), e) // 5 5 [1 2 3 4 5]

	f := make([]int, len(e))
	copy(f, e)
	fmt.Println("copy: ", f) // copy:  [1 2 3 4 5]

	l := d[2:5]
	fmt.Println(l, cap(l))

	//切片再切片
	a := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a:%v type:%T len:%d  cap:%d\n", a, a, len(a), cap(a))
	b := a[1:3]
	fmt.Printf("b:%v type:%T len:%d  cap:%d\n", b, b, len(b), cap(b))
	c := b[1:5]
	fmt.Printf("c:%v type:%T len:%d  cap:%d\n", c, c, len(c), cap(c))

}
