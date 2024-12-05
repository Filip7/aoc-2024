package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(rules *map[int][]int, pages *[][]int, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstPart := true
	for scanner.Scan() {
		var num1, num2 int
		if scanner.Text() == "" || scanner.Text() == "\n" {
			firstPart = false
			continue
		}
		if firstPart {
			fmt.Sscanf(scanner.Text(), "%d|%d", &num1, &num2)
			(*rules)[num1] = append((*rules)[num1], num2)
		} else {
			line := strings.Split(scanner.Text(), ",")
			var numLine []int
			for _, el := range line {
				num, _ := strconv.Atoi(el)
				numLine = append(numLine, num)
			}
			*pages = append(*pages, numLine)
		}
	}
}

func validate(curPos int, page *[]int, rule *[]int) bool {
	isStartOk, isEndOk := true, true
	for i := range *page {
		searchEl := (*page)[i]
		if searchEl == (*page)[curPos] {
			continue
		}

		if len(*rule) == 0 {
			break
		}

		idx := slices.IndexFunc(*rule, func(s int) bool { return s == searchEl })
		if i < curPos {
			isStartOk = idx == -1
			if !isStartOk {
				break
			}
		} else if i > curPos {
			isEndOk = idx != -1
			if !isEndOk {
				break
			}
		}
	}

	return isStartOk && isEndOk
}

func checkPages(rules *map[int][]int, pages *[][]int) (int, int) {
	sum := 0
	sumInc := 0
	correctPages := make([][]int, 0)
	incorrectPages := make([][]int, 0)
	pSize := len(*pages)

	for i := 0; i < pSize; i++ {
		shouldAdd := false
		page := (*pages)[i]
		for j, el := range page {
			rule := (*rules)[el]

			shouldAdd = validate(j, &page, &rule)
			if !shouldAdd {
				shouldAdd = false
				break
			}
		}
		if shouldAdd {
			correctPages = append(correctPages, page)
		} else {
			incorrectPages = append(incorrectPages, page)
		}
	}

	for _, page := range correctPages {
		mid := len(page) / 2
		sum += page[mid]
	}

	for _, page := range incorrectPages {
		slices.SortFunc(page, func(i, j int) int {
			idx := slices.IndexFunc((*rules)[i], func(s int) bool { return s == j })
			return idx
		})
		mid := len(page) / 2
		sumInc += page[mid]
	}

	return sum, sumInc
}

func main() {
	fmt.Println("AOC 2024 - DAY 5")

	rules := make(map[int][]int)
	pages := make([][]int, 0)
	readFile(&rules, &pages, "day5/data.txt")
	sum, sumInc := checkPages(&rules, &pages)

	fmt.Println("Sum is: ", sum, " sum of incorrect: ", sumInc)
}
