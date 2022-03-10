package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 2
	var num2 string = "2"

	var num1str string = strconv.Itoa(num1)
	num2int, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num2int == num1)
	}

	num3 := "1100"
	num3parse, err := strconv.ParseInt(num3, 2, 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num3parse)
	}

	fmt.Println(num1str == num2)
}
