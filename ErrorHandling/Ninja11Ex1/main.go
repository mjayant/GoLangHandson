package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	p1 := person{Name: "jayant",
		Age:    34,
		Speaks: []string{"abc", "def", "ghi"},
	}
	data, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("data is ", data)
	fmt.Printf("Type of data is %T", data)
}

type person struct {
	Name   string
	Age    int
	Speaks []string
}
