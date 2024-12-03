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

func parseFile() string {
	var lines []string

	file, err := os.Open("../example2.txt")
	// file, err := os.Open("../example.txt")
	// file, err := os.Open("../data.txt")
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

func handleAllMulInputs(line string) {
	res := 0
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)

	match := r.FindAllString(line, len(line))

	for _, str := range match {
		var num1, num2 int

		fmt.Sscanf(str, "mul(%d,%d)", &num1, &num2)

		res += num1 * num2
	}

	fmt.Println("All mul: ", res)
}

// This works only for example2 and data file
func handleDoAndDontMulInputs(line string) {
	res := 0

	rDoAndDont := regexp.MustCompile(`^(.*?)don't\(\)|do\(\)(.*?)don't\(\)|do\(\)(.*)`)
	rMul := regexp.MustCompile(`mul\(\d+,\d+\)`)

	matchM := rDoAndDont.FindAllString(line, -1)
	allInOneString := strings.Join(matchM, "")
	match := rMul.FindAllString(allInOneString, -1)

	for _, str := range match {
		var num1, num2 int

		fmt.Sscanf(str, "mul(%d,%d)", &num1, &num2)

		res += num1 * num2
	}

	fmt.Println("Only do(): ", res)
}

func main() {
	fmt.Println("AOC 2024 - DAY 3")

	line := parseFile()
	handleAllMulInputs(line)
	handleDoAndDontMulInputs(line)
}
