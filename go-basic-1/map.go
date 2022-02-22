package main

import "fmt"

func main() {
	person := make(map[string]string) // inisialisasi map kosong
	person2 := map[string]string{
		"name": "NooBee",
		"job":  "Programmer",
	} // inisialisasi map dengan langsung mendefinisikan key dan value

	computer := map[string]int{
		"sold":  14,
		"stock": 20,
	} // inisialisasi map dengan key dan value yang berbeda tipe data

	fmt.Println(person)
	fmt.Println(person2)
	fmt.Println(computer)
	fmt.Println("")

	person["name"] = "Reyhan"
	person["job"] = "Designer"

	person2["address"] = "Pekanbaru"

	computer["created_year"] = 2019

	fmt.Println(person)
	fmt.Println(person2)
	fmt.Println(computer)
}
