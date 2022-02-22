package main

import "fmt"

type Student struct {
	Name  string
	Major string
	Grade int
}

func (s *Student) SayHello() {
	cetak(s.Name, "sedang berada pada tingkat", s.Grade, "dan menempuh jurusan", s.Major)
}

func (s Student) SetName(name string) {
	s.Name = name
}
func (s *Student) SetNamePointer(name string) {
	s.Name = name
}

func cetak(arg ...interface{}) {
	fmt.Println(arg...)
}

func main() {
	var student Student
	student = Student{
		Name:  "NooBee",
		Major: "Teknik Informatika",
		Grade: 3,
	}

	student.SayHello()
	student.SetName("Nongky")
	student.SayHello()
	student.SetNamePointer("Nongky")
	student.SayHello()

}
