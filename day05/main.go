package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	"regexp"
	// "sort"
	"strconv"
	// "strings"
)

// read file and return numbers as []int
func readInput(path string) ([][4]int, error) {
	var nums [][4]int
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	re := regexp.MustCompile(`,| -> `)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		string_array := re.Split(line, -1)
		int_array := [4]int{}

		for i, num := range string_array {
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("strconv.Atoi: %s", err)
			}
			int_array[i] = val
		}

		nums = append(nums, int_array)
	}

	return nums, scanner.Err()
}

func main() {
	vectors, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readInput: %s", err)
	}

	// using maps lets me track only the points with lines on them
	// this is the only interesting choice in this whole script but you can keep reading anyway
	field_pt1 := make(map[[2]int]int)
	field_pt2 := make(map[[2]int]int)

	for _, vector := range vectors {

		// horizontal lines
		if vector[0] == vector[2] {

			// x1 = x2 and y1 < y2
			if vector[1] < vector[3] {
				for i := vector[1]; i <= vector[3]; i++ {
					field_pt1[[2]int{vector[0], i}]++
					field_pt2[[2]int{vector[0], i}]++
				}

			// x1 = x2 and y1 > y2
			} else {
				for i := vector[3]; i <= vector[1]; i++ {
					field_pt1[[2]int{vector[0], i}]++
					field_pt2[[2]int{vector[0], i}]++
				}
			}

		// vertical lines
		} else if vector[1] == vector[3] {

			// x1 < x2 and y1 = y2
			if vector[0] < vector[2] {
				for i := vector[0]; i <= vector[2]; i++ {
					field_pt1[[2]int{i, vector[1]}]++
					field_pt2[[2]int{i, vector[1]}]++
				}

			// x1 > x2 and y1 = y2
			} else {
				for i := vector[2]; i <= vector[0]; i++ {
					field_pt1[[2]int{i, vector[1]}]++
					field_pt2[[2]int{i, vector[1]}]++
				}
			}

		// diagonal lines
		} else if vector[0] < vector[2] {

			// x1 < x2 and y1 < y2
			if vector[1] < vector[3] {
				for i := 0; i <= vector[2] - vector[0]; i++ {
					field_pt2[[2]int{vector[0] + i, vector[1] + i}]++
				}

			// x1 < x2 and y1 > y2
			} else {
				for i := 0; i <= vector[2] - vector[0]; i++ {
					field_pt2[[2]int{vector[0] + i, vector[1] - i}]++
				}

			}
		} else {

			// x1 > x2 and y1 < y2
			if vector[1] < vector[3] {
				for i := 0; i <= vector[0] - vector[2]; i++ {
					field_pt2[[2]int{vector[0] - i, vector[1] + i}]++
				}

			// x1 > x2 and y1 > y2
			} else {
				for i := 0; i <= vector[0] - vector[2]; i++ {
					field_pt2[[2]int{vector[0] - i, vector[1] - i}]++
				}

			}
		}
	}

	total := 0
	for _, value := range field_pt1 {
		if value > 1 {
			total++
		}
	}
	fmt.Println("PART 1:", total)
	total = 0
	for _, value := range field_pt2 {
		if value > 1 {
			total++
		}
	}
	fmt.Println("PART 2:", total)
}
