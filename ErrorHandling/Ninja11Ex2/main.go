package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := person{
		name: "Jayant",
		age:  34,
	}
	data, err := toJson(p1)
	if err != nil {
		fmt.Println("Something went wrong", err)
	}
	fmt.Println("Json data is ,", data)
}

type person struct {
	name string
	age  int
}

func toJson(p person) ([]byte, error) {
	byte_data, err := json.Marshal(p)
	if err != nil {
		return []byte{}, fmt.Errorf("Something went wrong during marshling json %v", err)
	}
	return byte_data, nil
}
