package main

import "fmt"

func main() {
	fmt.Println(faktorial(5))
}

func fibo(size int) []int {
	var arr []int

	for i := 0; i < size; i++ {
		if len(arr) == 0 || len(arr) == 1 {
			arr = append(arr, 1)
		} else {
			arr = append(arr, arr[i-1]+arr[i-2])
		}
	}
	return arr
}

// latihan
// bilangan prima
func prime(size int) []int {
	totalLooping := 0
	var arr []int
	for i := 1; i <= size; i++ {
		counter := 0
		for j := 1; j <= size; j++ {
			// lakukan perulangan ..
			// counter = 3 => perulangan ke 11

			// if counter >= 3 {
			// 	break
			// }

			// execute
			if i%j == 0 {
				counter++
			}

			// counter = 3 => perulangan 10

			// looping stop
			if counter >= 3 {
				break
			}

			// counter => 3

			totalLooping++
		}

		if counter == 2 && i != 1 {
			arr = append(arr, i)
		}
	}

	fmt.Println(totalLooping)

	return arr
}

// bilangan faktorial
func faktorial(size int) int {
	result := 1

	for i := 2; i <= size; i++ {
		result *= i
	}

	return result
}
