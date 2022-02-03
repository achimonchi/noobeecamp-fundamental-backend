package main

import "fmt"

func main() {
	num := 2
	if num > 1 {
		text := "Lebih besar dari 1"
		fmt.Println(text)

	} else {
		text := "Lebih kecil atau sama dengan 1"
		fmt.Println(text)
		fmt.Println(num)
	}
}
