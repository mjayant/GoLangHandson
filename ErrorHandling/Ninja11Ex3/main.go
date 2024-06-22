package main

import "fmt"

func main() {
	ce1 := CustomError{
		CustomError: "Something went wrong",
	}
	foo(ce1)
}

type CustomError struct {
	CustomError string
}

func (ce CustomError) Error() string {
	return ce.CustomError
}

func foo(e error) {
	fmt.Println(e.Error())
}
