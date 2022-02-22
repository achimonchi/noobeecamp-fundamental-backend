package main

import "fmt"

func main() {
	arr := []int{10, 3, 8, 6, 1, 3, 2, 7}
	min := getMin(arr)
	max := getMax(arr)

	isExists := searchIndexAngka(arr, 80)
	fmt.Println(min, max, isExists)
}

func getMin(arr []int) int {
	min := 0

	min = arr[0]

	for _, a := range arr {
		if a < min {
			min = a
		}
	}

	return min

}

func getMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]

	for _, a := range arr {
		if a > max {
			max = a
		}
	}

	return max
}

func searchIndexAngka(arr []int, num int) int {
	index := -1

	for i, a := range arr {
		if a == num {
			index = i
			break
		}
	}

	return index
}
