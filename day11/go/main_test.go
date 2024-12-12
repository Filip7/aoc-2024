package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// func Test(t *testing.T) {
// 	res, expected := 0, 55312
// 	input := readFile("../example.txt")
//
// 	res = calc(input)
//
// 	if res != expected {
// 		t.Errorf("Result is %d and expected %d!", res, expected)
// 	}
// }

func Test3(t *testing.T) {
	str := readFile("../example.txt")
	strSpl := strings.Split(str, " ")
	input := map[int]int{}
	for _, el := range strSpl {
		num, _ := strconv.Atoi(el)
		input[num] = 1
	}

	steps := 3
	total := ReduceAndSum(steps, input)

	fmt.Println(total)
}

// func Test2(t *testing.T) {
// 	res, expected := 0, 100
// 	readFile("../example.txt")
//
// 	if res != expected {
// 		t.Errorf("Result is %d and expected %d!", res, expected)
// 	}
// }
