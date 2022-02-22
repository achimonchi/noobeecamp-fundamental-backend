package main

import "fmt"

type Fruit struct {
	Name   string
	Weight int
}

func main() {
	var fruit1 Fruit
	fruit1.Name = "Apple"
	fruit1.Weight = 1

	fruit2 := Fruit{
		Name:   "Mango",
		Weight: 2,
	}

	var fruit3 = Fruit{
		Name:   "Banana",
		Weight: 1,
	}

	fruit4 := Fruit{"Lemon", 4}

	fmt.Printf("%v\n", fruit1)
	fmt.Println(fruit2)
	fmt.Println(fruit3)
	fmt.Println(fruit4)
}
