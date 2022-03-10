package main

import "fmt"

type Calculate interface {
	Divide() float32
	Multiply() float32
}

type Num struct {
	Num1 float32
	Num2 float32
}

func (n *Num) Divide() float32 {
	return n.Num1 / n.Num2
}
func (n *Num) Multiply() float32 {
	return n.Num1 * n.Num2
}

func main() {
	// num := NewNum(30, 10)
	num := Num{
		Num1: 30,
		Num2: 10,
	}
	fmt.Println(num.Divide())
	fmt.Println(num.Multiply())
}
