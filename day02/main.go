package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// read file and return lines as []int
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var aim int = 0
	var x int = 0
	var y_basic int = 0
	var y_aim int = 0
	for i := 0; i < len(lines); i++ {
		line := strings.Fields(lines[i])
		val, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("strconv.Atoi: %s", err)
		}

		switch line[0] {
		case "forward":
			x += val
			y_aim += val * aim
		case "up":
			y_basic -= val
			aim -= val
		case "down":
			y_basic += val
			aim += val
		}
	}

	fmt.Println("x =", x)
	fmt.Println("y_basic =", y_basic)
	fmt.Println("y_aim =", y_aim)
	fmt.Println("x * y_basic =", x * y_basic)
	fmt.Println("x * y_aim =", x * y_aim)
}
