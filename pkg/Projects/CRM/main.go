package main

import (
	"fmt"
	"src/pkg/Projects/CRM/view"
)

func main() {
	err := view.NewCustomerView().MainMenu()
	if err != nil {
		fmt.Println(err)
	}
}
