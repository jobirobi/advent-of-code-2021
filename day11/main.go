package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	// "regexp"
	// "sort"
	// "strconv"
	// "strings"
)

// read file and return numbers as [][]int
func readLines(path string) ([][]int, error) {
	var nums [][]int
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var int_array []int
		for _, num := range scanner.Text() {

			// rune = number + 48 ¯\_(ツ)_/¯
			int_array = append(int_array, int(num) - 48)
		}

		nums = append(nums, int_array)
	}

	return nums, scanner.Err()
}

func flash(octopi [][]int, flashes int, x int, y int) (ret_flashes int) {
	ret_flashes = flashes
	if octopi[y][x] >= 10 {
		ret_flashes++
		octopi[y][x] = 0

		for yy := y-1; yy <= y+1; yy++ {
			if yy >= 0 && yy < 10 {
				for xx := x-1; xx <= x+1; xx++ {
					if xx >= 0 && xx < 10 {
						if octopi[yy][xx] != 0 {
							octopi[yy][xx]++
							ret_flashes = flash(octopi, ret_flashes, xx, yy)
						}
					}
				}
			}
		}
	}
	return
}

func main() {
	octopi, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	flashes := 0
	i := 1
	for ; i <= 100; i++ {

		// first increment
		for y, line := range octopi {
			for x, _ := range line {
				octopi[y][x]++
			}
		}

		// flash
		for y, line := range octopi {
			for x, _ := range line {
				flashes = flash(octopi, flashes, x, y)
			}
		}
	}
	fmt.Println("PART 1:", flashes)

	i--
	flashes = 0
	for flashes < 100 {
		i++
		flashes = 0

		// first increment
		for y, line := range octopi {
			for x, _ := range line {
				octopi[y][x]++
			}
		}

		// flash
		for y, line := range octopi {
			for x, _ := range line {
				flashes = flash(octopi, flashes, x, y)
			}
		}
	}
	fmt.Println("PART 2:", i)
}
