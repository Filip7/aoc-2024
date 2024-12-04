package main

import (
	"testing"
)

func Test(t *testing.T) {
	line := parseFile("../example.txt")
	resAll := handleAllMulInputs(line)
	resDoAndDont := handleDoAndDontMulInputs(line)

	if resAll != 161 || resDoAndDont != 0 {
		t.Errorf("Result was incorrect, got resAll: %d, resDoAndDont: %d", resAll, resDoAndDont)
	}
}

func Test2(t *testing.T) {
	line := parseFile("../example2.txt")
	resAll := handleAllMulInputs(line)
	resDoAndDont := handleDoAndDontMulInputs(line)

	if resAll != 161 || resDoAndDont != 48 {
		t.Errorf("Result was incorrect, got resAll: %d, resDoAndDont: %d", resAll, resDoAndDont)
	}
}
