package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {
	tv := TV{Goods{"电视机 001", 5000.1}, Brand{"海尔", "山东"}}
	tv2 := TV2{&Goods{"电视机 002", 9999}, &Brand{"华为", "河南"}}
	fmt.Println("tv1", tv)
	fmt.Println("tv2", tv2)

	fmt.Println("tv2 detail:", *tv2.Goods, tv2.Brand)
}
