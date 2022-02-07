package main

import "fmt"

func main() {
	gender := "woman"
	age := 20

	var canWork bool

	if gender == "man" {
		if age >= 18 {
			canWork = true
		} else {
			canWork = false
		}
	} else {
		if age >= 20 {
			canWork = true
		} else {
			canWork = false
		}
	}

	if !canWork { // ini sama artinya dengan canWork == false
		fmt.Println("Saya tidak boleh bekerja")
	} else {
		fmt.Println("Saya boleh bekerja")
	}
}
