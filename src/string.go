package main

import (
	"fmt"
)

func getName() (firstName, secondName, LastName string) {
	return "first_name", "second_name", "last_name"
}

func main() {
	_, _, last_name := getName()
	fmt.Printf("%s\n", last_name)
	fmt.Printf("%#v\n", []byte("Hello,世界"))

	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: go run strings.go string")
	//	os.Exit(1)
	//}

	//str := os.Args[1]
	//
	//fmt.Printf("str is type %s\n", reflect.TypeOf(str))

}

func init() {
	var a = "hello world"
	print(a)
	println()
}
