package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func file(rules *map[int][]int, pages *[][]int, filePath string) {
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
	isGood1 := true
	isGood2 := true

	for k := range *page {
		sEl := (*page)[k]
		if sEl == (*page)[curPos] {
			continue
		}

		if len(*rule) == 0 {
			break
		}

		idx := slices.IndexFunc(*rule, func(s int) bool { return s == sEl })
		if k < curPos {
			if idx == -1 {
				isGood1 = true
			} else {
				isGood1 = false
				break
			}
		} else if k > curPos {
			if idx != -1 {
				isGood2 = true
			} else {
				isGood2 = false
				break
			}
		}
	}

	return isGood1 && isGood2
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
		sort.Slice(page, func(i, j int) bool {
			idx := slices.IndexFunc((*rules)[page[i]], func(s int) bool { return s == page[j] })
			return idx != -1
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
	file(&rules, &pages, "day5/data.txt")
	sum, sumInc := checkPages(&rules, &pages)

	fmt.Println("Sum is: ", sum, " sum of incorrect: ", sumInc)
}
