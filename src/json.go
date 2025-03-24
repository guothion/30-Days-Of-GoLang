package main

import (
	"encoding/json"
	"fmt"
)

func main() {}

type Student struct {
	ID     int8 `json:"id"`
	Gender string
	name   string
}

func init() {
	s1 := Student{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data)
}
