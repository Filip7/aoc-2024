package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
	"unicode"
)

var (
	xSize int
	ySize int
)

type Position struct {
	x int
	y int
}

func (pos Position) isOutOfBounds() bool {
	return (pos.x < 0 || pos.y < 0) || (pos.x >= xSize || pos.y >= ySize)
}

func (pos1 Position) distance(pos2 Position) (int, int) {
	dX, dY := math.Abs(float64(pos1.x)-float64(pos2.x)), math.Abs(float64(pos1.y)-float64(pos2.y))
	return int(dX), int(dY)
}

func (pos1 Position) equals(pos2 Position) bool {
	return pos1.x == pos2.x && pos1.y == pos2.y
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

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

func readFile(matrix *[][]string, charMap *map[rune][]Position, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	xPos := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for yPos, let := range scanner.Text() {
			if unicode.IsNumber(let) || unicode.IsLetter(let) {
				(*charMap)[let] = append((*charMap)[let], Position{xPos, yPos})
			}
		}
		t := strings.Split(scanner.Text(), "")
		*matrix = append(*matrix, t)
		xPos += 1
	}
}

func calc(charMap *map[rune][]Position) int {
	pos := 0
	allLocations := []Position{}
	for _, els := range *charMap {
		for i := 0; i < len(els); i++ {
			el1 := els[i]
			for j := i + 1; j < len(els); j++ {
				el2 := els[j]

				dX, dY := el1.distance(el2)
				lPos, hPos := calculatePos(el1, el2, dY, dX)

				setAntinode(lPos, &allLocations)
				setAntinode(hPos, &allLocations)
			}
		}
	}

	for range allLocations {
		pos += 1
	}

	return pos
}

func calcResonance(charMap *map[rune][]Position) int {
	pos := 0
	allLocations := []Position{}
	for _, els := range *charMap {
		for i := 0; i < len(els); i++ {
			el1 := els[i]
			for j := i + 1; j < len(els); j++ {
				el2 := els[j]

				if i == 0 && j == 1 {
					firstPos := slices.IndexFunc(allLocations, func(s Position) bool { return s.equals(el1) })
					if firstPos == -1 {
						allLocations = append(allLocations, el1)
					}
					secondPos := slices.IndexFunc(allLocations, func(s Position) bool { return s.equals(el2) })
					if secondPos == -1 {
						allLocations = append(allLocations, el2)
					}
				}

				dX, dY := el1.distance(el2)
				lPos, hPos := calculatePos(el1, el2, dY, dX)

				lp1, lp2 := el1, lPos
				hp1, hp2 := el2, hPos
				for !lp1.isOutOfBounds() || !lp2.isOutOfBounds() ||
					!hp1.isOutOfBounds() || !hp2.isOutOfBounds() {
					setAntinode(lp1, &allLocations)
					setAntinode(lp2, &allLocations)
					lp1, lp2 = calculatePos(lp1, lp2, dY, dX)

					setAntinode(hp1, &allLocations)
					setAntinode(hp2, &allLocations)
					hp1, hp2 = calculatePos(lp1, lp2, dY, dX)
				}
			}
		}
	}

	for range allLocations {
		pos += 1
	}

	return pos
}

func calculatePos(el1 Position, el2 Position, dY int, dX int) (Position, Position) {
	lpx, lpy, hpx, hpy := -1, -1, -1, -1
	if el1.y > el2.y {
		lpy = el1.y + dY
		hpy = el2.y - dY
	} else if el1.y < el2.y {
		lpy = el1.y - dY
		hpy = el2.y + dY
	} else {
		lpy = el1.y
		hpy = el2.y
	}

	if el1.x > el2.x {
		lpx = el1.x + dX
		hpx = el2.x - dX
	} else if el1.x < el2.x {
		lpx = el1.x - dX
		hpx = el2.x + dX
	} else {
		lpx = el1.x
		hpx = el2.x
	}

	lPos, hPos := Position{lpx, lpy}, Position{hpx, hpy}
	return lPos, hPos
}

func setAntinode(pos Position, allLocations *[]Position) {
	if !pos.isOutOfBounds() {
		idx := slices.IndexFunc((*allLocations), func(s Position) bool { return s.equals(pos) })
		if idx == -1 {
			*allLocations = append((*allLocations), pos)
		}
	}
}

func main() {
	fmt.Println("AOC 2024 - DAY 8")
	matrix, charMap := [][]string{}, map[rune][]Position{}
	readFile(&matrix, &charMap, "day8/data.txt")
	xSize, ySize = len(matrix), len(matrix[0])
	positions := calc(&charMap)
	positionsResonance := calcResonance(&charMap)

	fmt.Println("Number of antinodes is: ", positions)
	fmt.Println("Resonance num of antinodes is: ", positionsResonance)
}
