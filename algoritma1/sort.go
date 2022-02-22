package main

import (
	"fmt"
	"math/rand"
)

func main() {
	generateRandomData := func(size int) []int {
		fmt.Printf("Generate random data with size %d \n", size)
		var arr []int
		counter := 1
		for counter <= size {
			arr = append(arr, rand.Intn(size))
			counter++
		}

		return arr
	}
	nums := generateRandomData(100)
	fmt.Println(nums)
	sort := insertionSort(nums)

	fmt.Println(sort)
}

func bubbleSort(arr []int) []int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
			count++
		}
	}

	fmt.Println(count)

	return arr
}

func insertionSort(arr []int) []int {
	count := 0
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
			count++
		}
	}
	fmt.Println(count)
	return arr
}
