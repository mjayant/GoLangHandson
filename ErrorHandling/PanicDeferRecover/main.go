package main

import "fmt"

func main() {
	fmt.Println("About to call function which have panic")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered panic from calling function. %v", r)
		}
	}()
	fmt.Println("Calling foo function")
	foo()
	fmt.Println("After recover panic, this line not suppose to run")
}

func foo() {
	fmt.Println("inside foo function")
	fmt.Println("About to call panic function")
	panic("Somethig went wrong during calling foo")
	fmt.Println("This line under foo function, won't run")
}
