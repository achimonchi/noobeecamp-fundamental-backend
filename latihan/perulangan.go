package main

import "fmt"

func main() {
	// soal 1
	/**
	*****
	****
	***
	**
	*
	 */
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	// soal 2
	/**
	*
	**
	***
	****
	*****
	 */

	fmt.Println("")

	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("")
	fmt.Println("")

	for i := 1; i <= 50; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("Noo Bee")
		} else if i%3 == 0 {
			fmt.Print("Noo")
		} else if i%5 == 0 {
			fmt.Print("Bee")
		} else {
			fmt.Print(i)
		}
		fmt.Print(", ")
	}
	fmt.Println("")
}
