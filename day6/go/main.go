package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Position struct {
	x int
	y int
}

func (pos1 Position) equal(pos2 Position) bool {
	return pos1.x == pos2.x && pos1.y == pos2.y
}

type Direction int

const (
	N = iota
	E
	S
	W
)

func readFile(matrix *[][]string, filePath string) Position {
	var guardPos Position
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	xPos, found := 0, false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), "")
		*matrix = append(*matrix, t)
		if !found {
			yPos := slices.Index(t, "^")
			if yPos != -1 {
				guardPos = Position{xPos, yPos}
				found = true
			}
			xPos += 1
		}
	}

	return guardPos
}

func getNextPos(direction Direction, pos Position) Position {
	nextX, nextY := -1, -1
	switch direction {
	case N:
		nextX, nextY = pos.x-1, pos.y
	case E:
		nextX, nextY = pos.x, pos.y+1
	case S:
		nextX, nextY = pos.x+1, pos.y
	case W:
		nextX, nextY = pos.x, pos.y-1
	}
	return Position{nextX, nextY}
}

func changeDir(curDir Direction) Direction {
	switch curDir {
	case N:
		return E
	case E:
		return S
	case S:
		return W
	case W:
		return N
	}

	return -1
}

func isOutOfBounds(pos Position, xSize int, ySize int) bool {
	return (pos.x < 0 || pos.y < 0) || (pos.x >= xSize || pos.y >= ySize)
}

func countSteps(matrix *[][]string, firstPos Position) int {
	walked := map[Position]bool{}
	walked[firstPos] = true
	direction := Direction(N)
	xSize, ySize := len(*matrix), len((*matrix)[0])

	curPos := firstPos
	for !isOutOfBounds(curPos, xSize, ySize) {
		nextPos := getNextPos(direction, curPos)
		if isOutOfBounds(nextPos, xSize, ySize) {
			break
		}
		nextChar := (*matrix)[nextPos.x][nextPos.y]
		if nextChar == "#" {
			direction = changeDir(direction)
			continue
		}

		walked[nextPos] = true
		curPos = nextPos
	}

	return len(walked)
}

func searchForLoop(matrix *[][]string, firstPos Position, obstacle Position) bool {
	walked := map[Position]bool{}
	walked[firstPos] = true
	direction := Direction(N)
	xSize, ySize := len(*(matrix)), len((*matrix)[0])

	been := map[Position][]Direction{}

	curPos := firstPos
	for !isOutOfBounds(curPos, xSize, ySize) {
		nextPos := getNextPos(direction, curPos)
		if isOutOfBounds(nextPos, xSize, ySize) {
			break
		}

		nextChar := (*matrix)[nextPos.x][nextPos.y]
		if nextChar == "#" || nextPos.equal(obstacle) {
			direction = changeDir(direction)
			continue
		}

		hasBeen := slices.IndexFunc(been[curPos], func(d Direction) bool {
			return d == direction
		})
		if hasBeen != -1 {
			return true
		} else {
			been[curPos] = append(been[curPos], direction)
		}

		walked[nextPos] = true
		curPos = nextPos
	}

	return false
}

func coutLoops(matrix *[][]string, firstPos Position) int {
	numOfLoops := 0
	xSize, ySize := len((*matrix)), len((*matrix)[0])

	sumNumOfLoops := map[Position]bool{}
	direction := Direction(N)
	sent := map[Position]bool{}

	curPos := firstPos
	for !isOutOfBounds(curPos, xSize, ySize) {
		nextPos := getNextPos(direction, curPos)
		if isOutOfBounds(nextPos, xSize, ySize) {
			break
		}
		nextChar := (*matrix)[nextPos.x][nextPos.y]
		if nextChar == "#" {
			direction = changeDir(direction)
			continue
		}

		if !sent[nextPos] {
			looped := searchForLoop(matrix, firstPos, nextPos)
			if looped {
				sumNumOfLoops[nextPos] = true
				sent[nextPos] = true
			}
		}

		curPos = nextPos
	}

	for range sumNumOfLoops {
		numOfLoops += 1
	}

	return numOfLoops
}

func main() {
	fmt.Println("AOC 2024 - DAY 6")

	matrix := [][]string{}
	guardPos := readFile(&matrix, "day6/data.txt")
	walked := countSteps(&matrix, guardPos)
	loops := coutLoops(&matrix, guardPos)

	fmt.Println("Guard will be on this many positions: ", walked)
	fmt.Println("This many loops positions there are: ", loops)
}
