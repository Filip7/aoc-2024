package main

import "testing"

func Test(t *testing.T) {
	res, expected := 0, 100
	readFile("../example.txt")

	if res != expected {
		t.Errorf("Result is %d and expected %d!", res, expected)
	}
}

func Test2(t *testing.T) {
	res, expected := 0, 100
	readFile("../example.txt")

	if res != expected {
		t.Errorf("Result is %d and expected %d!", res, expected)
	}
}
