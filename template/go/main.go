package main

import (
	"fmt"
	"log"
	"os"
)

func file(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func main() {
	fmt.Println("AOC 2024 - DAY ")
}
