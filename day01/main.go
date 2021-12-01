package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	"strconv"
)

// read file and return lines as []int
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, value)
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var monocount int = 0
	var tricount int = 0
	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i - 1] {
			monocount++
		}
		if i >= 3 && (lines[i - 2] + lines[i - 1] + lines[i]) > (lines[i - 3] + lines[i - 2] + lines[i - 1]) {
			tricount++
		}
	}
	fmt.Println("The number of one-step increases is", monocount)
	fmt.Println("The number of three-measurement increases is", tricount)
}
