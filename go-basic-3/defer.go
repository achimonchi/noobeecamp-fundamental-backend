package main

import (
	"fmt"
)

func main() {
	num1 := 5

	if num1%2 > 0 {
		cetak("Ini adalah bilangan ganjil")

		func() {
			cetak("akan berakhir setelah proses diatas selesai")
			cetak("Test")
		}()

	}

	cetak("Oh tentu tidak. Defer akan di eksekusi setelah ini !")
}

func cetak(text string) {
	fmt.Println(text)
}
