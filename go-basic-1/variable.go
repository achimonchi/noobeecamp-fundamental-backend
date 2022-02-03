package main

import (
	"fmt"
)

var age int = 22

const (
	DB   = "POSTGRES"
	HOST = "localhost"
	PORT = "5432"
)

func main() {
	ok, result, isSuccess := calculate()

	fmt.Println(ok, result, isSuccess, DB)
}

func calculate() (string, int, bool) {
	fmt.Println(HOST)
	return "oke", 100, true
}
