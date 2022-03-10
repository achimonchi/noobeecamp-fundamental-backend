package code2

import (
	"testing"
)

func TestBallsDefault(t *testing.T) {
	balls := []string{"merah", "biru", "biru", "kuning", "merah"}
	colors := []string{"merah", "kuning", "biru"}
	expected := map[string]int{
		"merah": 2,
		"biru":  2,
	}

	total := GetBallsByColor(balls, colors)

	if len(expected) != len(total) {
		t.Errorf("salah ! diharapkan warna yang muncul ada %d, namun yang didapat adalah %d", len(expected), len(total))
	}

	for color, count := range total {
		if expected[color] != count {
			t.Errorf("salah ! hasil yang diharapkan untuk warna - %s adalah :%d sedangkan yang didapat adalah :%d", color, expected[color], count)
		}
	}
}

func TestBallsWithCustomColors(t *testing.T) {
	arrBalls := map[int][]string{
		0: {"merah", "biru", "biru", "kuning", "merah"},
		1: {"merah", "biru", "kuning", "merah"},
		2: {"merah", "biru", "biru", "hijau", "merah", "hijau"},
		3: {"biru", "biru", "kuning", "merah", "biru"},
		4: {"kuning", "biru", "biru", "kuning", "merah"},
	}

	arrColors := map[int][]string{
		0: {"merah", "kuning", "biru"},
		1: {"merah", "kuning", "biru"},
		2: {"merah", "hijau", "biru"},
		3: {"merah", "kuning", "biru"},
		4: {"merah", "kuning", "biru"},
	}

	arrExpected := map[int]map[string]int{
		0: {
			"merah": 2,
			"biru":  2,
		},
		1: {
			"merah": 2,
		},
		2: {
			"merah": 2,
			"biru":  2,
			"hijau": 2,
		},
		3: {
			"biru": 3,
		},
		4: {
			"kuning": 2,
			"biru":   2,
		},
	}

	for key, ball := range arrBalls {
		expected := arrExpected[key]

		got := GetBallsByColor(ball, arrColors[key])

		if len(expected) != len(got) {
			t.Errorf("salah ! diharapkan warna yang muncul ada %d, namun yang didapat adalah %d", len(expected), len(got))
		}

		for color, count := range got {
			if expected[color] != count {
				t.Errorf("salah! hasil yang diharapkan untuk warna - %s pada index - %d adalah :%d sedangkan yang didapat adalah :%d", color, key, expected[color], count)
			}
		}
	}
}
