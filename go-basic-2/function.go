package main

import "fmt"

func main() {
	nilai := []int{5, 2, 8, 6, 10, 8, 9, 5, 8}
	kkm := 6

	print(graduate, nilai, kkm)
	// total, nilaiLulus := graduate(nilai, kkm)
}

func graduate(nilaiArr []int, limit int) (int, func() []int) {
	var studentsGraduate []int

	for _, arr := range nilaiArr {
		if arr >= limit {
			studentsGraduate = append(studentsGraduate, arr)
		}
	}

	nilaiLulus := func() []int {
		return studentsGraduate
	}

	return len(studentsGraduate), nilaiLulus
}

func print(g func(nilaiArr []int, limit int) (int, func() []int), nilai []int, kkm int) {
	total, nilaiLulus := g(nilai, kkm)
	fmt.Println(total, nilaiLulus())
}
