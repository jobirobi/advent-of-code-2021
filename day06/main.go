package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// read file and return numbers as []int
func readInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// only one line to scan in
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}
	num_strings := strings.Split(line, ",")

	// now convert to ints
	var nums []int
	for _, num := range num_strings {
		val, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("strconv.Atoi: %s", err)
		}
		nums = append(nums, val)
	}

	return nums, scanner.Err()
}

func count(nums []int) (counts [9]int) {
	for _, num := range nums {
		counts[num]++
	}
	return
}

func sum(counts [9]int) (total int) {
	for _, num := range counts {
		total += num
	}
	return
}

func main() {
	nums, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	counts := count(nums)

	i := 1
	for ; i <= 80; i++ {
		for num, count := range counts {
			if num == 0 {
				counts[0] -= count
				counts[6] += count
				counts[8] += count
			} else {
				counts[num] -= count
				counts[num - 1] += count
			}
		}
	}
	fmt.Println("After 80 days, there are", sum(counts), "lanternfish")

	for ; i <= 256; i++ {
		for num, count := range counts {
			if num == 0 {
				counts[0] -= count
				counts[6] += count
				counts[8] += count
			} else {
				counts[num] -= count
				counts[num - 1] += count
			}
		}
	}
	fmt.Println("After 256 days, there are", sum(counts), "lanternfish")
}
