package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 2, 3, 0, 5}
	// printArray(arr)
	fmt.Println(getMin(arr))

	// fact := factorial(-5)
	// fmt.Println(fact)
}

func printArray(arr []int) {
	fmt.Println("After :", arr)
	if len(arr) == 0 {
		return
	} else {
		fmt.Println(arr[0])
		fmt.Println("Before :", arr)
		printArray(arr[1:])

	}
}

func getMin(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	} else {
		if arr[0] < arr[1] {
			temp := arr[1]
			arr[1] = arr[0]
			arr[0] = temp
		}
		return getMin(arr[1:])
	}
}

func factorial(num int) int {
	if num == 0 {
		return 1
	} else if num == 1 {
		return num
	} else if num < 0 {
		return num * factorial(num+1)
	} else {
		return num * factorial(num-1)
	}
}
