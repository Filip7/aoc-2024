package main

import (
	"sort"
	"testing"
)

func Test(t *testing.T) {
	var col1, col2 []int
	simil := make(map[int]int)

	initColumns(&col1, &col2, "../example.txt")
	sort.Ints(col1)
	sort.Ints(col2)

	similarityScore(&col2, &simil)

	finalSum, sumSilim := totalDistanceAndSimilarity(&col1, &col2, &simil)

	if finalSum != 11 || sumSilim != 31 {
		t.Errorf("Result was incorrect, got finalSum: %d, sumSilim: %d", finalSum, sumSilim)
	}
}
