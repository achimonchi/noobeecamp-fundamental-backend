package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Name   string
	Height int
}

func (s *Student) Validate() error {
	// jika salah satu ada yang tidak diisi, kita tau ini akan menghasilkan error.
	// Nah untuk menghandle itu, kita akan membuat error kita sendiri
	// inilah yang dimaksud dengan error yang expected
	if s.Name == "" {
		return errors.New("Nama tidak boleh kosong")
	}

	if len(s.Name) <= 3 {
		panic(errors.New("Nama harus lebih besar dari 3"))
		// return errors.New("Nama harus lebih besar dari 3")
	}

	if s.Height == 0 {
		return errors.New("Tinggi tidak boleh kosong")
	}

	return nil
}

func main() {
	defer catchError()
	student := Student{Name: "NooBee", Height: 165}

	err := student.Validate()
	if err != nil {
		cetak(err.Error())
	} else {
		cetak("Hello, " + student.Name)
	}

	getData(cetak)
	getData(func(nama string) {
		fmt.Println(nama)
	})
}

func cetak(text string) {
	fmt.Println(text)
}

func print(text string) {
	fmt.Println("text")
}

func catchError() {
	err := recover()
	if err != nil {
		cetak("Ada error nih :" + fmt.Sprintf("%v", err))
	}
}

func getData(callback func(string)) {
	callback("Hello dari callback")
}
