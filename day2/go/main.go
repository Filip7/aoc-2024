package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func handleInput(sliceOfSlices *[][]int) {
	// file, err := os.Open("../example.txt")
	file, err := os.Open("../data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels := scanner.Text()
		levelsStringSlice := strings.Split(levels, " ")

		levelsIntSlice := make([]int, len(levelsStringSlice))
		for i, el := range levelsStringSlice {
			num, err := strconv.Atoi(el)
			if err != nil {
				log.Fatal(err)
			}
			levelsIntSlice[i] = num
		}
		*sliceOfSlices = append(*sliceOfSlices, levelsIntSlice)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func safeLevelsCountNoDampener(sliceOfSlices *[][]int) int {
	var plus bool
	positive := 0

	for _, level := range *sliceOfSlices {
		for iLvl := range level {
			jLvl := iLvl + 1
			localPlus := false
			if iLvl == len(level)-1 {
				positive += 1
				break
			}
			firstLvl := level[iLvl]
			secondLvl := level[jLvl]

			result := firstLvl - secondLvl
			if result > 0 {
				localPlus = true
			} else if result < 0 {
				localPlus = false
			} else {
				break
			}
			if iLvl == 0 {
				plus = localPlus
			}

			if localPlus != plus {
				break
			}

			absoluteResult := math.Abs(float64(result))

			if absoluteResult < 1 || absoluteResult > 3 {
				break
			}
		}
	}

	return positive
}

func safeLevelsCountWITHDampener(sliceOfSlices *[][]int) int {
	positive := 0

	for _, level := range *sliceOfSlices {
		positive += calculateSafeLvlDampener(&level, false)
	}

	return positive
}

func calculateSafeLvlDampener(level *[]int, dampenerActive bool) int {
	var plus bool
	for iLvl := range *level {
		jLvl := iLvl + 1
		localPlus := false

		if iLvl == len(*level)-1 {
			return 1
		}
		firstLvl := (*level)[iLvl]
		secondLvl := (*level)[jLvl]

		result := firstLvl - secondLvl
		if result > 0 {
			localPlus = true
		} else if result < 0 {
			localPlus = false
		} else {
			if !dampenerActive {
				return dampenerCalc(level, iLvl, jLvl)
			}
			break
		}
		if iLvl == 0 {
			plus = localPlus
		}

		if localPlus != plus {
			if !dampenerActive {
				return dampenerCalc(level, iLvl, jLvl)
			}
			break
		}

		absoluteResult := math.Abs(float64(result))

		if absoluteResult < 1 || absoluteResult > 3 {
			if !dampenerActive {
				return dampenerCalc(level, iLvl, jLvl)
			}
			break
		}
	}
	return 0
}

func dampenerCalc(level *[]int, iLvl int, jLvl int) int {
	shouldReturn1, returnValue1 := calc(level, iLvl, jLvl)
	if shouldReturn1 {
		return returnValue1
	}

	shouldReturn2, returnValue2 := calc(level, iLvl+1, jLvl+1)
	if shouldReturn2 {
		return returnValue2
	}

	if iLvl-1 >= 0 {
		shouldReturn3, returnValue3 := calc(level, iLvl-1, jLvl-1)
		if shouldReturn3 {
			return returnValue3
		}
	}

	return 0
}

func calc(level *[]int, i int, j int) (bool, int) {
	slice := make([]int, len(*level))
	copy(slice, *level)
	slice = slices.Delete(slice, i, j)
	res := calculateSafeLvlDampener(&slice, true)
	if res == 1 {
		return true, res
	}
	return false, 0
}

func main() {
	fmt.Println("AOC 2024 - DAY 2")
	sliceOfSlices := make([][]int, 0)
	handleInput(&sliceOfSlices)

	safeLevels := safeLevelsCountNoDampener(&sliceOfSlices)
	safeLevelsDampener := safeLevelsCountWITHDampener(&sliceOfSlices)

	fmt.Println("Safe levels with NO dampener: ", safeLevels)
	fmt.Println("Safe levels WITH dampener: ", safeLevelsDampener)
}
