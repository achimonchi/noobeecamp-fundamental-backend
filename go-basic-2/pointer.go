package main

import "fmt"

func main() {
	var num1 int = 5
	var num2 *int = &num1

	fmt.Println("===== Nilai Awal ===== ")
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	fmt.Println("===== Nilai Setelah num1 Diubah ===== ")
	num1 = 6
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	fmt.Println("===== Nilai Setelah num2 Diubah ===== ")
	*num2 = 10
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(*num2)
}
