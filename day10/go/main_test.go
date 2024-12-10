package main

import (
	"testing"
)

func Test(t *testing.T) {
	res, excpected := 0, 36
	matrix := [][]string{}
	startPos := []Position{}
	readFile(&matrix, &startPos, "../example.txt")

	res = hike(&matrix, &startPos, true)

	if res != excpected {
		t.Errorf("Got %d and expected %d!", res, excpected)
	}
}

func Test2(t *testing.T) {
	res, excpected := 0, 81
	matrix := [][]string{}
	startPos := []Position{}
	readFile(&matrix, &startPos, "../example.txt")

	res = hike(&matrix, &startPos, false)

	if res != excpected {
		t.Errorf("Got %d and expected %d!", res, excpected)
	}
}
