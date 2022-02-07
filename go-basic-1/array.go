package main

import "fmt"

func main() {
	// horizonal
	// fruits := [3]string{"Apple", "Banana", "Cherry"}

	// vertical
	// fruits := [3]string{
	// 	"Apple",
	// 	"Banana",
	// 	"Cherry",
	// }

	// make
	// fruits := make([]string, 3)

	// fruits[0] = "Apple"
	// fruits[1] = "Banana"
	// fruits[2] = "Cherry"

	var fruits = [2][3]string{
		{"Apple", "Mango", "Banana"},
		{"Orange", "Grape", "Avocado"},
	}

	fmt.Println(fruits)
	fmt.Println(fruits[0])
	fmt.Println(fruits[0][1])
	fmt.Println(fruits[1])
	fmt.Println(fruits[1][2])

	// fmt.Println(fruits)
}
