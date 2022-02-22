package main

import "fmt"

func main() {
	animals := []string{"Ayam", "Sapi", "Garuda"}

	anim1 := animals[0:1]
	anim2 := animals[1:2]

	fmt.Println("====[ANIMALS]====")
	fmt.Println("Data \t :", animals)
	fmt.Println("Capacity :", cap(animals))
	fmt.Println("Len \t :", len(animals))
	fmt.Println("")

	fmt.Println("====[ANIM1]====")
	fmt.Println("Data \t :", anim1)
	fmt.Println("Capacity :", cap(anim1))
	fmt.Println("Len \t :", len(anim1))
	fmt.Println("")

	fmt.Println("====[ANIM2]====")
	fmt.Println("Data \t :", anim2)
	fmt.Println("Capacity :", cap(anim2))
	fmt.Println("Len \t :", len(anim2))
	fmt.Println("")

	fmt.Println("Append ke anim 1")
	anim1 = append(anim1, "Gajah")
	fmt.Println("====[ANIMALS]====")
	fmt.Println("Data \t :", animals)
	fmt.Println("Capacity :", cap(animals))
	fmt.Println("Len \t :", len(animals))
	fmt.Println("")

	fmt.Println("====[ANIM1]====")
	fmt.Println("Data \t :", anim1)
	fmt.Println("Capacity :", cap(anim1))
	fmt.Println("Len \t :", len(anim1))
	fmt.Println("")

	fmt.Println("====[ANIM2]====")
	fmt.Println("Data \t :", anim2)
	fmt.Println("Capacity :", cap(anim2))
	fmt.Println("Len \t :", len(anim2))
	fmt.Println("")

	fmt.Println("Append ke anim 2")
	anim2 = append(anim2, "Unta")
	fmt.Println("====[ANIMALS]====")
	fmt.Println("Data \t :", animals)
	fmt.Println("Capacity :", cap(animals))
	fmt.Println("Len \t :", len(animals))
	fmt.Println("")

	fmt.Println("====[ANIM1]====")
	fmt.Println("Data \t :", anim1)
	fmt.Println("Capacity :", cap(anim1))
	fmt.Println("Len \t :", len(anim1))
	fmt.Println("")

	fmt.Println("====[ANIM2]====")
	fmt.Println("Data \t :", anim2)
	fmt.Println("Capacity :", cap(anim2))
	fmt.Println("Len \t :", len(anim2))
	fmt.Println("")

	fmt.Println("Append ke anim 1")
	anim1 = append(anim1, "Monyet")
	fmt.Println("====[ANIMALS]====")
	fmt.Println("Data \t :", animals)
	fmt.Println("Capacity :", cap(animals))
	fmt.Println("Len \t :", len(animals))
	fmt.Println("")

	fmt.Println("====[ANIM1]====")
	fmt.Println("Data \t :", anim1)
	fmt.Println("Capacity :", cap(anim1))
	fmt.Println("Len \t :", len(anim1))
	fmt.Println("")

	fmt.Println("====[ANIM2]====")
	fmt.Println("Data \t :", anim2)
	fmt.Println("Capacity :", cap(anim2))
	fmt.Println("Len \t :", len(anim2))
	fmt.Println("")
}
