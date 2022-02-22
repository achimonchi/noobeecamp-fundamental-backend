package main

import "fmt"

func main() {
	// nilai := map[string]int{
	// 	"Agama":      80,
	// 	"Matematika": 90,
	// 	"Olahraga":   70,
	// 	"Design":     80,
	// }

	// fmt.Println("===== Nilai Awal ======")
	// fmt.Println(len(nilai))
	// fmt.Println(nilai)
	// fmt.Println("\n\n===== Nilai Akhir ======")
	// delete(nilai, "Design")
	// fmt.Println(len(nilai))
	// fmt.Println(nilai)

	// key := "Agama"
	// val, isExists := nilai[key]

	// if isExists {
	// 	fmt.Println(val)
	// } else {
	// 	fmt.Println("Data", key, "tidak ditemukan")
	// }

	students := []map[string]string{
		{"name": "NooBee", "major": "Teknik Komputer"},
		{"name": "Jovie", "major": "Sistem Informasi"},
		{"name": "Reyhan", "major": "Teknik Informatika"},
	}

	for _, student := range students {
		fmt.Println(student["name"], "berada di jurusan", student["major"])
	}
}
