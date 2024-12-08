package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func readFile(eqs *map[int][]int, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numLine := strings.Split(scanner.Text(), ":")
		res, _ := strconv.Atoi(strings.TrimSpace(numLine[0]))
		argsString := strings.TrimSpace(numLine[1])
		args := strings.Split(argsString, " ")
		for _, arg := range args {
			num, _ := strconv.Atoi(string(arg))
			(*eqs)[res] = append((*eqs)[res], num)
		}
	}
}

func generateOperands(numOfOperands int, checkCombineOperator bool) []string {
	os := []string{"+", "*"}
	if checkCombineOperator {
		os = append(os, "|")
	}
	os2 := os

	for range numOfOperands - 1 {
		tmp := []string{}
		for _, o := range os {
			for _, o2 := range os2 {
				tmp = append(tmp, o2+o)
			}
		}
		os2 = tmp
	}

	return os2
}

func calculate(elems []int, operands string, c chan int, excpetedRes int) {
	defer wg.Done()
	result := elems[0]
	oS := strings.Split(operands, "")

	for i := 0; i < len(operands); i++ {
		switch oS[i] {
		case "*":
			result *= elems[i+1]
		case "+":
			result += elems[i+1]
		case "|":
			result, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(elems[i+1]))
		}
		if result > excpetedRes {
			return
		}
	}

	c <- result
}

func calculateNonGo(elems []int, operands string, excpetedRes int) int {
	result := elems[0]
	oS := strings.Split(operands, "")

	for i := 0; i < len(operands); i++ {
		switch oS[i] {
		case "*":
			result *= elems[i+1]
		case "+":
			result += elems[i+1]
		case "|":
			result, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(elems[i+1]))
		}
		if result > excpetedRes {
			return result
		}
	}

	return result
}

func checkIfPossible(eqs *map[int][]int, checkCombineOperator bool) int {
	sumMe := sync.Map{}
	for res, eq := range *eqs {
		c := make(chan int)
		possibleOperands := len(eq) - 1
		operands := generateOperands(possibleOperands, checkCombineOperator)
		for _, o := range operands {
			wg.Add(1)
			go func() {
				calculate(eq, o, c, res)
			}()
		}

		go func() {
			wg.Wait()
			close(c)
		}()

		for el := range c {
			if el == res {
				sumMe.Store(res, true)
			}
		}
	}

	possible := 0
	sumMe.Range(func(key, val interface{}) bool {
		possible += key.(int)
		return true
	})

	return possible
}

func checkIfPossibleNonGo(eqs *map[int][]int, checkCombineOperator bool) int {
	sumMe := map[int]bool{}
	for res, eq := range *eqs {
		possibleOperands := len(eq) - 1
		operands := generateOperands(possibleOperands, checkCombineOperator)
		for _, o := range operands {
			calc := calculateNonGo(eq, o, res)
			if calc == res {
				sumMe[res] = true
				break
			}
		}
	}

	possible := 0
	for el := range sumMe {
		possible += el
	}

	return possible
}

func main() {
	fmt.Println("AOC 2024 - DAY 7")

	eqs := map[int][]int{}
	readFile(&eqs, "day7/data.txt")
	// resPlusAndMul := checkIfPossible(&eqs, false)
	// resAll := checkIfPossible(&eqs, true)

	resPlusAndMul := checkIfPossibleNonGo(&eqs, false)
	resAll := checkIfPossibleNonGo(&eqs, true)

	fmt.Println("Sum of possible combinations with + *: ", resPlusAndMul)
	fmt.Println("Sum of possible combinations with + * |: ", resAll)
}
