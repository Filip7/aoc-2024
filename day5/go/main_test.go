package main

import (
	"testing"
)

func Test(t *testing.T) {
	rules := make(map[int][]int)
	pages := make([][]int, 0)
	readFile(&rules, &pages, "../example.txt")

	sum, sumInc := checkPages(&rules, &pages)

	if sum != 143 {
		t.Errorf("Result was incorrect, got %d", sum)
	}

	if sumInc != 123 {
		t.Errorf("Corrected result was incorrect, got %d", sumInc)
	}
}
