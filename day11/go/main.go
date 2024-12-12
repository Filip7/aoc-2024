// package main
//
// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )
//
// func readFile(filePath string) string {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
//
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		return scanner.Text()
// 	}
//
// 	return ""
// }
//
// func walk(stone string) int {
// 	newStones := []string{stone}
// 	for i := range 75 {
// 		if i < 10 {
// 			fmt.Println("Stone: ", stone, " newStones: ", newStones)
// 		}
// 		tmpStones := []string{}
// 		for _, stn := range newStones {
// 			tStone := strings.Trim(stn, " ")
// 			if tStone == "0" {
// 				tmpStones = append(tmpStones, "1")
// 			} else if len(tStone)%2 == 0 {
// 				half := (len(tStone) / 2)
// 				new1, new2 := tStone[:half], tStone[half:]
// 				num1, _ := strconv.Atoi(new1)
// 				num2, _ := strconv.Atoi(new2)
//
// 				tmpStones = append(tmpStones, strconv.Itoa(num1))
// 				tmpStones = append(tmpStones, strconv.Itoa(num2))
// 			} else {
// 				num, _ := strconv.Atoi(tStone)
// 				num *= 2024
// 				tmpStones = append(tmpStones, strconv.Itoa(num))
// 			}
// 		}
// 		newStones = tmpStones
// 	}
// 	return len(newStones)
// }
//
// func calc(input string) int {
// 	numStones := 0
// 	newInput := input
// 	stones := strings.Split(newInput, " ")
// 	for _, stone := range stones {
// 		numStones += walk(stone)
// 	}
//
// 	return numStones
// }
//
// func main() {
// 	fmt.Println("AOC 2024 - DAY 11")
// 	input := readFile("day11/data.txt")
//
// 	res := calc(input)
//
// 	fmt.Println("Result is: ", res)
// }

package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func blink(input map[int]int) map[int]int {
	result := make(map[int]int)
	for k, v := range input {
		result = change(k, v, result)
	}
	return result
}

func change(n int, count int, result map[int]int) map[int]int {
	if n == 0 {
		result[1] += count
		return result
	}

	length := len(strconv.Itoa(n))
	if length%2 == 0 {
		half := int(math.Pow(10, float64(length/2)))
		left := n / half
		right := n % half

		result[left] += count
		result[right] += count
	} else {
		result[n*2024] += count
	}
	return result
}

func ReduceAndSum(steps int, stones map[int]int) int {
	acc := stones
	for range steps {
		acc = blink(acc)
	}

	sum := 0
	for _, el := range acc {
		sum += el
	}
	return sum
}

func readFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func main() {
	// Example usage
	str := readFile("day11/data.txt")
	strSpl := strings.Split(str, " ")
	input := map[int]int{}
	for _, el := range strSpl {
		num, _ := strconv.Atoi(el)
		input[num] = 1
	}

	steps := 2
	total := ReduceAndSum(steps, input)
	println("Total sum:", total)
}
