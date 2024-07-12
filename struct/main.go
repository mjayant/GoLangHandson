package main

import "fmt"

func main() {

	type Person struct {
		name           string
		first_name     string
		fav_ice_creams []string
	}

	p1 := Person{
		name:           "abc",
		first_name:     "firstname_Abc",
		fav_ice_creams: []string{"a1", "a2", "b3"},
	}
	p2 := Person{
		name:           "xyz",
		first_name:     "firstname_xyz",
		fav_ice_creams: []string{"b1", "b2", "b3"},
	}
	a1 := []Person{p1, p2}
	for _, item := range a1 {
		fmt.Println("name is", item.name)
		fmt.Println("first name is", item.first_name)
		ice_creams := item.fav_ice_creams
		for _, ice_cream := range ice_creams {
			fmt.Printf("%v fav ice cream is %v \n", item.name, ice_cream)
		}

	}
	type person struct {
		first_name     string
		last_name      string
		fav_ice_creams []string
	}
	m1 := map[string]person{}
	p11 := person{
		last_name:      "xyz",
		first_name:     "firstname_xyz",
		fav_ice_creams: []string{"b1", "b2", "b3"},
	}
	p22 := person{
		last_name:      "abc",
		first_name:     "firstname_abc",
		fav_ice_creams: []string{"a1", "a2", "a3"},
	}
	m1[p1.first_name] = p11
	m1[p2.first_name] = p22
	for key, val := range m1 {
		fmt.Println("printing value of map")
		fmt.Println("printing value of map key", key)
		fmt.Println("printing value of map val", val)
		fmt.Println("Printing value of struct")
		fmt.Println("firs name struct", val.first_name)
		fmt.Println("last name struct", val.last_name)
		fmt.Println("fav ice creams ", val.fav_ice_creams)

	}
	type Engine struct {
		electric bool
	}
	type Vechile struct {
		engine Engine
		model  string
		doors  int
		color  string
	}
	e1 := Engine{
		electric: true,
	}
	e2 := Engine{
		electric: false,
	}
	v1 := Vechile{
		engine: e1,
		model:  "V1",
		doors:  4,
		color:  "grey",
	}
	v2 := Vechile{
		engine: e2,
		model:  "V2",
		doors:  4,
		color:  "Red",
	}
	p111 := struct {
		name      string
		friends   map[string]int
		favDrinks []string
	}{
		name: "james",
		friends: map[string]int{
			"james friends1": 45,
			"james friends2": 34,
		},
		favDrinks: []string{"d1", "d2", "d3"},
	}

	fmt.Println(p111)
	fmt.Printf("Vechile name is %s and engine is electric or not %t \n", v1.model, v1.engine.electric)
	fmt.Printf("Vechile name is %s and engine is electric or not %t \n", v2.model, v2.engine.electric)
}
