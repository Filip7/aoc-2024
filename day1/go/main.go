package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func initColumns(col1 *[]int, col2 *[]int, simil *map[int]int) {
	file, err := os.Open("../example.txt")
	// file, err := os.Open("../data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var num1 int
		var num2 int

		fmt.Sscanf(scanner.Text(), "%d   %d", &num1, &num2)
		*col1 = append(*col1, num1)
		*col2 = append(*col2, num2)
		(*simil)[num1] = 0
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func totalDistance(col1 *[]int, col2 *[]int) int {
	sum := 0
	for i := range *col1 {
		var max, min int
		num1 := (*col1)[i]
		num2 := (*col2)[i]
		if num1 > num2 {
			max = num1
			min = num2
		} else {
			max = num2
			min = num1
		}

		sum += max - min
	}

	return sum
}

func similarityScore(col2 *[]int, simil *map[int]int) {
	for _, el := range *col2 {
		(*simil)[el] += 1
	}
}

func printColumns(col1 *[]int, col2 *[]int) {
	for i := range *col1 {
		fmt.Println((*col1)[i], "   ", (*col2)[i])
	}
}

func main() {
	fmt.Println("AOC 2024 - DAY 1")

	var col1 []int
	var col2 []int
	simil := make(map[int]int)

	initColumns(&col1, &col2, &simil)

	sort.Slice(col1, func(i, j int) bool {
		return col1[i] < col1[j]
	})

	sort.Slice(col2, func(i, j int) bool {
		return col2[i] < col2[j]
	})

	// printColumns(&col1, &col2)

	finalSum := totalDistance(&col1, &col2)

	similarityScore(&col2, &simil)
	var sumSilim uint64 = 0
	for _, elem := range col1 {
		val, exists := simil[elem]
		if exists {
			sumSilim += uint64(elem * val)
		}
	}

	fmt.Println("Final sum: ", finalSum)
	fmt.Println("Similarity score: ", sumSilim)
}
