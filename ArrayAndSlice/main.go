package main

import "fmt"

func main() {
	x := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	fmt.Println(x)
	for index, item := range x {
		fmt.Println("index is ", index)
		fmt.Println("item is ", item)
	}

	var x1 [6]string
	x1 = [6]string{"abc", "def", "sg", "Sfgsf"}
	fmt.Println(x1)

	a1 := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	for _, item := range a1 {
		fmt.Println("value item ", item)
	}
	fmt.Println("length of slice", len(a1))

	a2 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a2); i++ {
		a2[i] = i + 1
	}
	for _, item := range a2 {
		fmt.Printf("value of item is %v and type is %T \n", item, item)
	}
	new_val := 8
	a3 := [10]int{}
	// copy(a3[:], a2)
	// var a2 [6]int
	for index, item := range a2 {
		a3[index] = item
	}
	a2[3] = new_val
	fmt.Printf("array is %v and type is %T\n", a2, a2)
	fmt.Printf("array is %v and type is %T", a1, a1)
	// a1 := []int{1, 2, 3, 4, 56, 57, 6898, 9, 987, 98, 9, 8}

	// i := 0
	// for i < len(a1) {
	// 	fmt.Printf("value is %v and type is %T \n", a1[i], a1[i])
	// 	i++
	// }
	//x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// a1 := x[:5]
	// a2 := x[5:]
	// a3 := x[2:7]
	// a4 := x[1:6]
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)
	// fmt.Println(a4)
	// x = append(x, 52)
	// fmt.Println(x)
	// x = append(x, 53, 54, 55)
	// fmt.Println("after adding 53, 54, 55 \t", x)
	// y := []int{56, 57, 58, 59, 60}
	// fmt.Println(" slice Y", y)
	// x = append(x, y...)
	// fmt.Println("after adding y slice to x")
	// fmt.Println(x)
	// x = append(x[:3], x[6:]...)
	// fmt.Println("After deleting values")
	// fmt.Println(x)
	// x := make([]string, 0, 60)
	// x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	// fmt.Println("slice length ", len(x))
	// fmt.Println("Slice capicity ", cap(x))
	// i := 0
	// for i < len(x) {
	// 	fmt.Println("value is \t", x[i])
	// 	fmt.Println("Index is \t", i)
	// 	i++
	// }
	aa1 := []string{"aa", "bb", "cc", "dd"}
	aa2 := []string{"a", "b", "c", "d", "e", "f"}
	var aa3 [][]string
	aa3 = append(aa3, aa1)
	aa3 = append(aa3, aa2)
	for index, val := range aa3 {
		fmt.Println("outer loop")
		fmt.Println("index is ", index)
		fmt.Println("value is", val)
		fmt.Println("------------outer loop ends here----------")
		for i, item := range val {
			fmt.Println("Iner loop")
			fmt.Println("index is ", i)
			fmt.Println("value is", item)
			fmt.Println("*********************Inner loop ends here*******************")
		}
	}

}
