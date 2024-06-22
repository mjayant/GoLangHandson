/*
Problem statement : write function sqrt , if it recive any negative value then
return cutomize error
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	val, err := sqRoot(2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(val)
}
func sqRoot(a float64) (float64, error) {
	if a < 0 {
		return 0.0, SqureRootCustomError{info: "Value is negative"}
	}
	return math.Sqrt(a), nil
}

type SqureRootCustomError struct {
	info string
}

func (sqrtE SqureRootCustomError) Error() string {
	return fmt.Sprintf(" Something went wrong during sqrt %v", sqrtE.info)
}
