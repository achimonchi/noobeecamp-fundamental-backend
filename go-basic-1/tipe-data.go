package main

import "fmt"

func main() {
	var teks string = `Hallo "NooBee"
Umur saya 14 Tahun
	`

	var age uint8 = 0
	var bigNumber int64 = -9223372036854775808

	fmt.Println(teks)
	fmt.Println(age)
	fmt.Println(bigNumber)
}
