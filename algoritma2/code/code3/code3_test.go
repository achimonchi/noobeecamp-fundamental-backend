package code3

import (
	"fmt"
	"testing"
)

func TestShortestAtoE(t *testing.T) {
	path := map[string]map[string]int{
		"A": {
			"B": 2,
			"C": 2,
		},
		"B": {
			"A": 2,
			"D": 2,
		},
		"C": {
			"A": 2,
			"D": 2,
		},
		"D": {
			"E": 1,
			"B": 2,
		},
		"E": {
			"D": 1,
			"F": 1,
		},
		"F": {
			"C": 3,
			"E": 1,
			"G": 2,
		},
		"G": {
			"E": 1,
			"F": 2,
			"H": 1,
		},
		"H": {
			"G": 1,
		},
	}

	from := "A"
	to := "E"

	expected := map[int]map[string]interface{}{
		0: {
			"path":  "A, B, D, E",
			"jarak": 5,
		},
		1: {
			"path":  "A, C, D, E",
			"jarak": 5,
		},
	}

	shortestPath := ShortestPath(from, to, path)
	fmt.Println(shortestPath, expected)
}
