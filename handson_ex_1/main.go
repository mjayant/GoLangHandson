package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Iota example
	const (
		a = iota
		b
		c
	)
	var d int = 45
	fmt.Printf("value of a is %d , b %d , c %d, d %d", a, b, c, d)

	/*
		Problem statement :
		Hands-on exercise #10 (was #05) - zero value, :=, type specificity,
		blank identifier
		write a program that uses the following:
		● var for zero value
		● short declaration operator
		● multiple initializations
		● var when specificity is required
		● blank identifier
		print items as necessary to make the program interesting
	*/

	for index, _ := range [10]string{} {
		fmt.Println(index)
	}
	var a2 int
	fmt.Println(" A2 values housl be zero", a2)
	a1, a4, a3 := 1, 2, 3
	fmt.Printf("Value of a1 is %d , a4 %d and a3 %d", a1, a4, a3)

	// Exercise for random number and string
	for i := 0; i <= 100; i++ {
		fmt.Println("Randome number:", rand.Intn(100))
	}
	var min int = 500
	var max int = 1000
	fmt.Println("Random number from 500 to 1000")
	for i := 0; i <= 100; i++ {
		fmt.Println("Randome number:", min+rand.Intn(max-min))
	}

	p1 := Person{
		name: "Mishra",
		age:  34,
	}
	p2 := Person{
		name: "Prerna",
		age:  32,
	}
	e1 := Employee{
		Person:     p1,
		department: "Development",
	}
	e2 := Employee{
		Person:     p2,
		department: "HR",
	}
	e1.getName()
	e2.getName()
}

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	department string
}

func (p1 Person) getName() {
	fmt.Println("Person name is ", p1.name)
}
