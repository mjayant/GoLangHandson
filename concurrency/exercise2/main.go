package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) speak() {
	fmt.Println(" Speak \t", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	p1 := Person{name: "name1",
		age: 34}
	p2 := Person{name: "name2",
		age: 3}
	saySomething(&p1)
	saySomething(&p2)
	p1.speak()

}
