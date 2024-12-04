package main

import (
	"testing"
)

func Test(t *testing.T) {
	sliceOfSlices := make([][]int, 0)
	handleInput(&sliceOfSlices, "../example.txt")

	safeLevels := safeLevelsCountNoDampener(&sliceOfSlices)
	safeLevelsDampener := safeLevelsCountWITHDampener(&sliceOfSlices)

	if safeLevels != 2 || safeLevelsDampener != 4 {
		t.Errorf("Result was incorrect, got safeLevels: %d, safeLevelsDampener: %d", safeLevels, safeLevelsDampener)
	}
}
