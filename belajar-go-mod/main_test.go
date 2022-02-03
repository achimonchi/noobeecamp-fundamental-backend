package main

import "testing"

func Calculate(a, b int) int {
	return a - b
}

func TestCalculate(t *testing.T) {
	a, b := 10, 20
	must := 30
	result := Calculate(a, b)

	t.Logf("Hasil Calculate adalah : %v", result)

	if result != must {
		t.Errorf("Salah ! Hasil yang diharapkan : %v dan yang didapatkan : %v", must, result)
	}
}
