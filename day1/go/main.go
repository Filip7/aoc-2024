package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func initColumns(col1 *[]int, col2 *[]int, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var num1, num2 int
		fmt.Sscanf(scanner.Text(), "%d   %d", &num1, &num2)

		*col1 = append(*col1, num1)
		*col2 = append(*col2, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func totalDistanceAndSimilarity(col1 *[]int, col2 *[]int, simil *map[int]int) (int, int) {
	totalDistance, similarity := 0, 0
	for i, elem := range *col1 {
		num1 := float64((*col1)[i])
		num2 := float64((*col2)[i])
		totalDistance += int(math.Abs(num1 - num2))

		val, exists := (*simil)[elem]
		if exists {
			similarity += elem * val
		}
	}

	return totalDistance, similarity
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

	var col1, col2 []int
	simil := make(map[int]int)

	initColumns(&col1, &col2, "day1/data.txt")
	sort.Ints(col1)
	sort.Ints(col2)

	similarityScore(&col2, &simil)

	finalSum, sumSilim := totalDistanceAndSimilarity(&col1, &col2, &simil)

	fmt.Println("Final sum: ", finalSum)
	fmt.Println("Similarity score: ", sumSilim)
}
