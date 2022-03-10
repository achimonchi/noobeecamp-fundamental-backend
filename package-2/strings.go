package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Bootcamp Backend NooBeeID"
	cari := "noobeeid"

	isExists := strings.Contains(text, cari)

	sliceText := strings.Split(text, " ")

	arrText := []string{"mangga", "jambu", "rambutan"}
	joinText := strings.Join(arrText, ", ")
	fmt.Println(isExists)
	fmt.Println(text)
	fmt.Println(sliceText[2])
	fmt.Println(joinText)
}
