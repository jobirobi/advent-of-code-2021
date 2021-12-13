package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	// "regexp"
	// "sort"
	"strconv"
	"strings"
)

// read file and return (xi,yi) as [][2]int and (x=xo, y=yo, etc) as []string
func readLines(path string) ([][2]int, []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var nums [][2]int
	var folds []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// skip empty lines
		if line == "" {
			continue
		}

		splits := strings.Split(line, ",")
		// "x,y" lines
		if len(splits) == 2 {
			xy := [2]int{0,0}
			for i, n := range splits {
				val, err := strconv.Atoi(n)
				if err != nil {
					return nil, nil, err
				}
				xy[i] = val
			}

			nums = append(nums, xy)

		// "fold along" lines
		} else {
			folds = append(folds, strings.Split(line, " ")[2])
		}
	}
	return nums, folds, scanner.Err()
}

func main() {
	nums, folds, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// using a map with dots represented as `true`
	field := make(map[[2]int]bool)
	for _, xy := range nums {
		field[xy] = true
	}

	// for each fold...
	for i, fold := range folds {

		// ...interpret the number part of "x=xo", "y=yo", then ...
		val, err := strconv.Atoi(strings.Split(fold, "=")[1])
		if err != nil {
			log.Fatalf("strconv.Atoi: %s", err)
		}

		// fold dots to the left
		if strings.HasPrefix(fold, "x") {
			for xy, _ := range field {
				if xy[0] > val{
					field[[2]int{xy[0] - 2 * (xy[0] - val), xy[1]}] = true
					delete(field, xy)
				}
			}

		// fold dots upward
		} else {
			for xy, _ := range field {
				if xy[1] > val{
					field[[2]int{xy[0], xy[1] - 2 * (xy[1] - val)}] = true
					delete(field, xy)
				}
			}
		}

		// after the first fold, count dots
		if i == 0 {
			fmt.Println("PART 1:", len(field))
		}
	}

	// a dumb but effective way to print out part 2:
	fmt.Println("PART 2:")
	maxx, maxy := 0,0
	for k, _ := range field {
		if k[0] > maxx {
			maxx = k[0]
		}
		if k[1] > maxy {
			maxy = k[1]
		}
	}
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			if field[[2]int{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
