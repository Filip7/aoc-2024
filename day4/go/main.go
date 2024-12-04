package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Direction int

const (
	N = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func parseFile(matrix *[][]string) {
	file, err := os.Open("day4/example.txt")
	// file, err := os.Open("day4/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), "")
		*matrix = append(*matrix, t)
	}
}

func getNextRowAndCol(direction Direction, curRow int, curCol int) (int, int) {
	nextRow, nextCol := -1, -1
	switch direction {
	case N:
		nextRow = curRow - 1
		nextCol = curCol
	case NE:
		nextRow = curRow - 1
		nextCol = curCol + 1
	case E:
		nextRow = curRow
		nextCol = curCol + 1
	case SE:
		nextRow = curRow + 1
		nextCol = curCol + 1
	case S:
		nextRow = curRow + 1
		nextCol = curCol
	case SW:
		nextRow = curRow + 1
		nextCol = curCol - 1
	case W:
		nextRow = curRow
		nextCol = curCol - 1
	case NW:
		nextRow = curRow - 1
		nextCol = curCol - 1
	}
	return nextRow, nextCol
}

func searchXmas(matrix *[][]string, curRow int, curCol int, searchedChar string, direction Direction) bool {
	numRow, numCol := len(*matrix), len((*matrix)[0])
	nextRow, nextCol := getNextRowAndCol(direction, curRow, curCol)

	if nextRow == -1 || nextCol == -1 || nextRow == numRow || nextCol == numCol {
		return false
	}

	nextChar := (*matrix)[nextRow][nextCol]
	if searchedChar == "M" && nextChar == searchedChar {
		return searchXmas(matrix, nextRow, nextCol, "A", direction)
	}
	if searchedChar == "A" && nextChar == searchedChar {
		return searchXmas(matrix, nextRow, nextCol, "S", direction)
	}
	if searchedChar == "S" && nextChar == searchedChar {
		return true
	}

	return false
}

func searchX_mas(matrix *[][]string, curRow int, curCol int) bool {
	rowNW, colNW := getNextRowAndCol(Direction(NW), curRow, curCol)
	rowSE, colSE := getNextRowAndCol(Direction(SE), curRow, curCol)

	rowNE, colNE := getNextRowAndCol(Direction(NE), curRow, curCol)
	rowSW, colSW := getNextRowAndCol(Direction(SW), curRow, curCol)

	nw, se := (*matrix)[rowNW][colNW], (*matrix)[rowSE][colSE]

	ne, sw := (*matrix)[rowNE][colNE], (*matrix)[rowSW][colSW]

	if (nw == "M" && se == "S" || nw == "S" && se == "M") &&
		(ne == "M" && sw == "S" || ne == "S" && sw == "M") {
		return true
	}

	return false
}

func searchXMAS(matrix *[][]string) (int, int) {
	numRow, numCol := len(*matrix), len((*matrix)[0])
	xmasCount, masCount := 0, 0

	for row := 0; row < numRow; row++ {
		for col := 0; col < numCol; col++ {
			if (*matrix)[row][col] == "X" {
				for i := 0; i <= 7; i++ {
					found := searchXmas(matrix, row, col, "M", Direction(i))
					if found {
						xmasCount++
					}
				}
			}
			if (row > 0 && row < numRow-1) && (col > 0 && col < numCol-1) {
				if (*matrix)[row][col] == "A" {
					found := searchX_mas(matrix, row, col)
					if found {
						masCount++
					}
				}
			}
		}
	}

	return xmasCount, masCount
}

func main() {
	fmt.Println("AOC 2024 - DAY 4")

	matrix := make([][]string, 0)
	parseFile(&matrix)
	xmasNum, masNum := searchXMAS(&matrix)

	fmt.Println("XMAS apears: ", xmasNum)
	fmt.Println("X-MAS apears: ", masNum)
}
