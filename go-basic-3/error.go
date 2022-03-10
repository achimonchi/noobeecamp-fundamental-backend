package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "10"

	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		cetak(numStr + " bukan sebuah angka")
		cetak(err.Error())
	} else {
		cetak(fmt.Sprintf("%d", num) + " adalah sebuah angka")
	}
}

func cetak(text string) {
	fmt.Println(text)
}
