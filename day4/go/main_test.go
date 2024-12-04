package main

import (
	"testing"
)

func Test(t *testing.T) {
	matrix := make([][]string, 0)
	parseFile(&matrix, "../example.txt")
	xmasNum, masNum := searchXMAS(&matrix)

	if xmasNum != 18 || masNum != 9 {
		t.Errorf("Result was incorrect, got xmasNum: %d, masNum: %d", xmasNum, masNum)
	}
}
