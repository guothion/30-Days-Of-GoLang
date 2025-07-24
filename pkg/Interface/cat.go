package main

import "fmt"

type Cat struct {
}

func (c Cat) say() string {
	return "喵喵喵"
}

type Dog struct{}

func (d Dog) say() string {
	return "汪汪汪"
}

func main() {
	c := Cat{}
	fmt.Println(c.say())
	d := Dog{}
	fmt.Println(d.say())
}
