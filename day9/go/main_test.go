package main

import (
	"testing"
)

func Test(t *testing.T) {
	res, excpected := 0, 1928
	dll := readFile("../example.txt")

	dll.SortPartition()
	dll.PrintForward()
	res = calc(dll)

	if res != excpected {
		t.Errorf("Got %d and excpected %d!\n", res, excpected)
	}
}

func Test2(t *testing.T) {
	res, excpected := 0, 2858
	dll := readFile("../example.txt")

	dll.SortCompact()
	dll.PrintForward()
	res = calc(dll)
	if res != excpected {
		t.Errorf("Got %d and excpected %d! <- partitioned\n", res, excpected)
	}
}
