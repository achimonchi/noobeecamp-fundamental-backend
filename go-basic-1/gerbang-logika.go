package main

import "fmt"

func main() {
	gender := "woman"
	age := 20

	var canWork bool

	if (gender == "man" && age >= 17) || (gender == "woman" && age >= 20) {
		canWork = true
	} else {
		canWork = false
	}

	if !canWork { // ini sama artinya dengan canWork == false
		fmt.Println("Saya tidak boleh bekerja")
	} else {
		fmt.Println("Saya boleh bekerja")
	}
}
