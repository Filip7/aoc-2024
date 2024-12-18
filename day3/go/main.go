package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// https://regex101.com/ <- life saver
func parseFile(filePath string) string {
	var lines []string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return strings.Join(lines, "")
}

func handleAllMulInputs(line string) int {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	match := r.FindAllString(line, -1)

	return caluclateResult(match)
}

// This works only for example2 and data file
func handleDoAndDontMulInputs(line string) int {
	rDoAndDont := regexp.MustCompile(`^(.*?)don't\(\)|do\(\)(.*?)don't\(\)|do\(\)(.*)`)
	rMul := regexp.MustCompile(`mul\(\d+,\d+\)`)

	matchM := rDoAndDont.FindAllString(line, -1)
	allInOneString := strings.Join(matchM, "")
	match := rMul.FindAllString(allInOneString, -1)

	return caluclateResult(match)
}

func caluclateResult(match []string) int {
	res := 0
	for _, str := range match {
		var num1, num2 int
		fmt.Sscanf(str, "mul(%d,%d)", &num1, &num2)

		res += num1 * num2
	}
	return res
}

func main() {
	fmt.Println("AOC 2024 - DAY 3")

	line := parseFile("day3/data.txt")
	resAll := handleAllMulInputs(line)
	resDoAndDont := handleDoAndDontMulInputs(line)

	fmt.Println("All mul: ", resAll)
	fmt.Println("Only do(): ", resDoAndDont)
}
