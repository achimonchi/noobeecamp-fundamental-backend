package code1

import (
	"testing"
)

func TestCandlesDefault(t *testing.T) {
	candles := []int{3, 2, 4, 4, 1}
	expected := 2

	count := GetCandles(candles)

	if count != expected {
		t.Errorf("salah ! hasil yang diharapkan adalah :%d sedangkan yang di dapat adalah :%d", expected, count)
	}
}

func TestCandlesWithCustomNumber(t *testing.T) {
	arrCandles := map[int][]int{
		0: {2, 8, 5, 2, 7, 3, 9, 4, 4, 1, 9, 7, 2},
		1: {1, 2, 3, 4, 5, 4, 2, 6, 4, 4, 1, 2, 4, 10, 7},
		2: {8, 4, 2, 9, 18, 5, 3, 18, 3, 17, 18, 9, 13},
		3: {9, 3, 7, 9, 2, 5, 1, 6, 2, 4, 9, 7, 9, 2, 3, 9},
		4: {8, 3, 5, 2, 8, 4, 5, 8, 2, 1, 8, 1, 7, 5, 2},
	}

	arrExpected := map[int]int{
		0: 2,
		1: 1,
		2: 3,
		3: 5,
		4: 4,
	}

	for key, candles := range arrCandles {
		expected := arrExpected[key]

		got := GetCandles(candles)
		if got != expected {
			t.Errorf("salah ! hasil yang diharapkan adalah :%d sedangkan yang di dapat adalah :%d", expected, got)
		}
	}
}

func TestCandleEmpty(t *testing.T) {
	candles := []int{}
	expected := 0

	count := GetCandles(candles)

	if count != expected {
		t.Errorf("salah ! hasil yang diharapkan adalah :%d sedangkan yang di dapat adalah :%d", expected, count)
	}
}
