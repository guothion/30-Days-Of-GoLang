package main

import "fmt"

func init() {
	s := "abnc"
	for i, c := range s {
		fmt.Print(i, c, "\t", string(c), "\n")
	}
}

func main() {}
