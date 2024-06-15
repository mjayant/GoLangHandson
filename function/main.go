package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width

}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)

}

type shape interface {
	area() float64
}

func info(a1 shape) float64 {
	a2 := a1.area()
	return a2
}
func main() {
	// Example for interface
	s1 := square{length: 12,
		width: 13,
	}
	fmt.Println("area of sqare is ", s1.area())
	c1 := circle{radius: 15}
	fmt.Println("Area of circle is ,", c1.area())
	fmt.Println("area of squre is,", info(s1))
	fmt.Println("area of circle is,", info(c1))

	// Example for , return value with mentioning in return statement , funtion at 135
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("result from foo", sum(a1))

	// Example for variadic param  func defination at 130
	a2 := foo(1, 2, 4, 5, 6, 67)
	fmt.Println(a2)

	//fmt.Println("result from foo", foo())
	//defer bar()

	//fmt.Println("Exit from main function")

	// Example for bind a method to struct
	p1 := Person{Name: "jayant",
		Age: 35}
	p1.speak()

	// Example for anonymous function , without calling inline
	x := func() {
		fmt.Println("Anonymous func ran")
	}

	y := func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}

	x()
	y("Todd")
	// func foo() int {
	// 	return 07
	// }

	// x()
	// y := func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(i)
	// 	}
	// }
	// y()
	// x := outer()
	// fmt.Println(x())
	// fmt.Println(printSqure(square, 5))
	// x := closureExample(2)
	// fmt.Println(x())
	// fmt.Println(x())
	// fmt.Println(x())
	// fmt.Println(x())
	// fmt.Println(x())
	// fmt.Println(x())
	// fmt.Println(x())
	// wrapperExample(sqre, 5)
	// var x *int
	// a1 := 5
	// x = &a1
	// fmt.Printf("value of x is %v \t and type %T \t and vaue stored at address is %v \n", x, x, *x)
	// var y *string
	// a2 := "jayant"
	// y = &a2
	// fmt.Printf("Value of y is %v \t and type is %T \t and value stored at adress is %v\n", y, y, *y)
	// p1 := person_details{name: "jayant",
	// 	age: 34}
	// p1.walk()
	// p1.run()
	// p1 = changeName(&p1, "Mishra")
	// p1.walk()
	// p1.run()

	// Example for unmarshalling JSON
	jsonData := `{"name":"Alice","age":30}`

	var p Person
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println(p.Name) // Output: {Alice 30}

}

func (p Person) speak() {
	fmt.Println("person age is ", p.Age)
	fmt.Println("person name is ", p.Name)

}

func foo(s1 ...int) int {
	fmt.Println("Inside the foo")
	total := 0
	for _, val := range s1 {
		total = total + val
	}
	return total
}

func sum(sl []int) (total int) {
	fmt.Println("Inside the sum")
	for _, val := range sl {
		total = total + val
	}
	return
}

// func bar() (int, string) {
// 	fmt.Println("Inside the bar")
// 	return 12, " return from bar"
// }

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// type person interface {
// 	walk()
// 	run()
// }
// type person_details struct {
// 	name string
// 	age  int
// }

// func (p person_details) walk() {
// 	fmt.Printf("%v name walk \n", p.name)
// }
// func (p person_details) run() {
// 	fmt.Printf("%v name run \n", p.name)
// }
// func changeName(p1 *person_details, s string) person_details {
// 	p1.name = s
// 	return *p1
// }

// func test(a *int) {
// 	fmt.Println("Inside function")
// 	fmt.Println("Value of a", a)
// 	fmt.Println(" Adress of param a", &a)
// 	*a = 56
// 	fmt.Println("After chaging value")
// 	fmt.Println("Value of a", a)
// }

// var x = func() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// func outer() func() int {
// 	return func() int {
// 		return 45
// 	}
// }

// func square(n int) int {
// 	return n * n
// }

// func printSqure(f func(n int) int, a int) string {
// 	return fmt.Sprintf("squre is %v", f(a))
// }

// func closureExample(val float64) func() float64 {
// 	var c float64
// 	return func() float64 {
// 		c++
// 		return math.Pow(val, c)
// 	}

// }
// func sqre(n int) int {
// 	return n * n
// }

// func wrapperExample(f func(n int) int, a int) {
// 	fmt.Println(" Before function")
// 	ret_value := f(a)
// 	fmt.Println("function return value", ret_value)
// 	fmt.Println("After calling function")
// }
