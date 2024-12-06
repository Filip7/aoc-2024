package main

import (
	"testing"
)

func Test(t *testing.T) {
	matrix := make([][]string, 0)
	guardPos := readFile(&matrix, "../example.txt")

	walked := countSteps(&matrix, guardPos)
	loops := coutLoops(&matrix, guardPos)

	if walked != 41 {
		t.Errorf("Result was incorrect, got %d walks", walked)
	}

	if loops != 6 {
		t.Errorf("Result was incorrect, got %d loops", loops)
	}
}
