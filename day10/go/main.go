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

var (
	xSize int
	ySize int
)

type Position struct {
	x int
	y int
}

func (direction Direction) getNextPos(pos Position) Position {
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

func (pos1 Position) equals(pos2 Position) bool {
	return pos1.x == pos2.x && pos1.y == pos2.y
}

func (pos Position) isOutOfBounds() bool {
	return (pos.x < 0 || pos.y < 0) || (pos.x >= xSize || pos.y >= ySize)
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

func readFile(matrix *[][]string, startPos *[]Position, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	x := 0
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), "")
		*matrix = append(*matrix, t)
		for y, el := range t {
			if el == "0" {
				*startPos = append(*startPos, Position{x, y})
			}
		}
		x += 1
	}
	xSize = x
	ySize = len((*matrix)[0])
}

func splitHike(matrix *[][]string, pos Position, curDif int, placesIveBeen *[]Position, unique bool) int {
	if pos.isOutOfBounds() || (*matrix)[pos.x][pos.y] != strconv.Itoa(curDif) {
		return 0
	}
	if (*matrix)[pos.x][pos.y] == "9" && curDif == 9 {
		if unique {
			firstPos := slices.IndexFunc(*placesIveBeen, func(s Position) bool { return s.equals(pos) })
			if firstPos == -1 {
				*placesIveBeen = append(*placesIveBeen, pos)
				return 1
			}
			return 0
		}
		return 1
	}

	return splitHike(matrix, N.getNextPos(pos), curDif+1, placesIveBeen, unique) +
		splitHike(matrix, E.getNextPos(pos), curDif+1, placesIveBeen, unique) +
		splitHike(matrix, W.getNextPos(pos), curDif+1, placesIveBeen, unique) +
		splitHike(matrix, S.getNextPos(pos), curDif+1, placesIveBeen, unique)
}

func hike(matrix *[][]string, startPos *[]Position, unique bool) int {
	sum := 0
	for _, pos := range *startPos {
		sum += splitHike(matrix, pos, 0, &[]Position{}, unique)
	}
	return sum
}

func main() {
	fmt.Println("AOC 2024 - DAY 10")
	matrix := [][]string{}
	startPos := []Position{}

	readFile(&matrix, &startPos, "day10/data.txt")

	sumUnique := hike(&matrix, &startPos, true)
	sumNonUnique := hike(&matrix, &startPos, false)

	fmt.Println("Unique result is: ", sumUnique)
	fmt.Println("Non unique result is: ", sumNonUnique)
}
