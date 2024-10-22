package main

import (
	"fmt"
	"os"
	"reflect"
)

func getName() (firstName, secondName, LastName string) {
	return "first_name", "second_name", "last_name"
}

func main() {
	_, _, last_name := getName()
	fmt.Printf("%s", last_name)

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run strings.go string")
		os.Exit(1)
	}

	str := os.Args[1]

	fmt.Printf("str is type %s\n", reflect.TypeOf(str))

}
