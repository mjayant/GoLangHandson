package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type userList []user

//	func (u userList) Len() int {
//		return len(u)
//	}
//
//	func (u userList) Swap(i, j int) {
//		u[i], u[j] = u[j], u[i]
//	}
//
//	func (u userList) Less(i, j int) bool {
//		return u[i].Age < u[j].Age
//	}
func main() {
	var Peopels []Person
	p1 := Person{Name: "name",
		Age: 12}
	p2 := Person{Name: "name1", Age: 4}
	p3 := Person{Name: "name2", Age: 56}
	Peopels = append(Peopels, p1, p2, p3)

	bs, err := json.Marshal(Peopels)
	if err != nil {
		fmt.Println("Somethign went wrong")
	}
	fmt.Println("---------------------------------")
	fmt.Println("Result of JSON marshal")
	fmt.Println(string(bs))
	fmt.Println("---------------------------------")
	s := `[{"Name":"name","Age":12},{"Name":"name1","Age":4},{"Name":"name2","Age":56}]`
	err = json.Unmarshal([]byte(s), &Peopels)
	if err != nil {
		fmt.Println("Somethign went wrong")
	}
	fmt.Println("**********************")
	fmt.Println("Result of JSON unmarshal")
	fmt.Println(Peopels)
	fmt.Println("**********************")

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)
	userList1 := userList(users)
	// sort.Sort(userList1)
	sort.Slice(userList1, func(i int, j int) bool { return userList1[i].Age < userList1[j].Age })
	fmt.Println("After sorting")
	fmt.Println(userList1)
	for _, item := range userList1 {
		fmt.Println("item is ------------>", item)
		saying := item.Sayings
		sort.Strings(saying)
		fmt.Println("sorted saying")
		fmt.Println(saying)
	}
	// err := json.NewEncoder(os.Stdout).Encode(users)
	// if err != nil {
	// 	fmt.Println("Something went wrong")
	// }
	// fmt.Println(users)
	// fmt.Printf("Type of output %T \n", users)
	// xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	// xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	// fmt.Println(xi)
	// sort.Ints(xi)
	// fmt.Println(xi)

	// fmt.Println(xs)
	// sort.Strings(xs)
	// fmt.Println(xs)
}
