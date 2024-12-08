package main

import (
	"testing"
)

func Test(t *testing.T) {
	expected, expectedR := 14, 34
	matrix := [][]string{}
	charMap := map[rune][]Position{}
	readFile(&matrix, &charMap, "../example.txt")
	xSize, ySize = len(matrix), len(matrix[0])
	count := calc(&charMap)
	positionsResonance := calcResonance(&charMap)

	if count != expected {
		t.Errorf("Wrong result got: %d and expected %d", count, expected)
	}
	if positionsResonance != expectedR {
		t.Errorf("Wrong result got: %d and expected %d", positionsResonance, expectedR)
	}
}
