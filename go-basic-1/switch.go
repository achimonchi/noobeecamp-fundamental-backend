package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Print(t.Hour())

	switch {
	case t.Hour() > 12:
		fmt.Println("PM")
	case t.Hour() < 12:
		fmt.Println("AM")
	}

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Working ...")
	// }

}
