package main

import (
	"fmt"
	"time"
)

func main() {
	// layout := "2006-01-02 15:04:05"
	myTimeStr := "2021-01-16T10:00:00Z"

	myTime, err := time.Parse(time.RFC3339, myTimeStr)
	if err != nil {
		fmt.Println("error :", err.Error())
	} else {
		fmt.Println("Waktu :", myTime)
		now := time.Now()
		nowFormat := now.Format(time.RubyDate)
		fmt.Println("Waktu sekarang:", now.String())
		fmt.Println("Tahun sekarang:", now.Year())
		fmt.Println("Bulan sekarang:", now.Month())
		fmt.Println("Tanggal sekarang:", now.Day())
		fmt.Println("Hari :", now.Weekday())
		fmt.Println(nowFormat)
	}

}
