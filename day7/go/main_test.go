package main

import (
	"testing"
)

func Test(t *testing.T) {
	eqs := map[int][]int{}
	readFile(&eqs, "../example.txt")
	resPlusAndMul := checkIfPossible(&eqs, false)
	resAll := checkIfPossible(&eqs, true)

	if resPlusAndMul != 3749 {
		t.Errorf("Calculation was wrong got: %d", resPlusAndMul)
	}
	if resAll != 11387 {
		t.Errorf("Calculation was wrong got: %d", resAll)
	}
}
