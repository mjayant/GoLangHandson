package main

import "fmt"

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
}
