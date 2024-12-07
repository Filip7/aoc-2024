package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	}
}

func main() {
	fmt.Println("AOC 2024 - DAY X")
	readFile("dayX/data.txt")
}
