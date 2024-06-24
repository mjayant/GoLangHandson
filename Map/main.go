package main

import (
	"fmt"
	"sort"
)

func main() {
	a1 := map[string][]string{}
	a1["bond_james"] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	a1["moneypenny_jenny"] = []string{`intelligence`, `literature`, `computer science`}
	a1["no_dr"] = []string{`cats`, `ice cream`, `sunsets`}

	delete(a1, "no_dr")
	if val, ok := a1["no_dr"]; !ok {
		fmt.Println("Key has been deleted")
	} else {
		fmt.Println("Key is not deleted , val is", val)
	}
	for key, val := range a1 {
		fmt.Println("------------outer------------loop")
		fmt.Println("Key is ", key)
		fmt.Println("Value is ", val)
		fmt.Println("--------------outer loop end here-------------")
		for index, item := range val {
			fmt.Println("**********************inner loop start here***************")
			fmt.Println("index:", index)
			fmt.Println("value : ", item)
			fmt.Println("**********************inner loop end here***************")
		}
	}
	//sort dict
	type map_struct struct {
		key   string
		value []string
	}
	var map_struct_lst []map_struct
	for k, v := range a1 {
		map_struct_lst = append(map_struct_lst, map_struct{k, v})
	}
	fmt.Println("Map is ", map_struct_lst)
	fmt.Printf("Type is %T \n", map_struct_lst)
	sort.Slice(map_struct_lst, func(i, j int) bool { return map_struct_lst[i].key > map_struct_lst[j].key })
	fmt.Println("After sorting")
	fmt.Println(map_struct_lst)
}
