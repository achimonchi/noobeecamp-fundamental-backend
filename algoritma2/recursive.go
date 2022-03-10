package main

import "fmt"

func main() {
	arr := []int{4, 2, 3, 0, 5}
	// printArray(arr)
	fmt.Println(getMin(arr))
}

func printArray(arr []int) {
	if len(arr) == 0 {
		return
	} else {
		fmt.Println(arr[0])
		printArray(arr[1:])
	}
}

func getMin(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	} else {
		fmt.Println(arr)
		if arr[0] < arr[1] {
			temp := arr[1]
			arr[1] = arr[0]
			arr[0] = temp
		}
		return getMin(arr[1:])
	}
}
