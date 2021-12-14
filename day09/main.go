package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	// "regexp"
	"sort"
	// "strconv"
	// "strings"
)

// read file and return numbers as [][]int
func readInput(path string) ([][]int, error) {
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

	return transpose(nums), scanner.Err()
}

func transpose(slice [][]int) [][]int {
    xl := len(slice[0])
    yl := len(slice)
    result := make([][]int, xl)
    for i := range result {
        result[i] = make([]int, yl)
    }
    for i := 0; i < xl; i++ {
        for j := 0; j < yl; j++ {
            result[i][j] = slice[j][i]
        }
    }
    return result
}

func explore(init [2]int, here [2]int, maxx int, maxy int, field [][]int, basin_sizes map[[2]int]int, explored map[[2]int]bool) {
	x, y := here[0], here[1]

	// look left:
	// if left exists, is unexplored, and is not a peak, explore left
	if x-1 >= 0 && ! explored[[2]int{x-1, y}] && field[x-1][y] != 9 {
		there := [2]int{x-1, y}
		basin_sizes[init]++
		explored[there] = true
		explore(init, there, maxx, maxy, field, basin_sizes, explored)
	}

	// look right:
	if x+1 <= maxx && ! explored[[2]int{x+1, y}] && field[x+1][y] != 9 {
		there := [2]int{x+1, y}
		basin_sizes[init]++
		explored[there] = true
		explore(init, there, maxx, maxy, field, basin_sizes, explored)
	}

	// look up:
	if y-1 >= 0 && ! explored[[2]int{x, y-1}] && field[x][y-1] != 9 {
		there := [2]int{x, y-1}
		basin_sizes[init]++
		explored[there] = true
		explore(init, there, maxx, maxy, field, basin_sizes, explored)
	}

	// look down:
	if y+1 <= maxy && ! explored[[2]int{x, y+1}] && field[x][y+1] != 9 {
		there := [2]int{x, y+1}
		basin_sizes[init]++
		explored[there] = true
		explore(init, there, maxx, maxy, field, basin_sizes, explored)
	}
}

func main() {
	field, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readInput: %s", err)
	}

	risk_level := 0
	maxx := len(field) - 1
	maxy := len(field[0]) - 1

	// I hate that golang forces me do this:
	basins := make([][]int, maxy+1)
	for i, _ := range basins {
		basins[i] = make([]int, maxx+1)
	}

	// find all the local minima
	explored := make(map[[2]int]bool)
	basin_sizes := make(map[[2]int]int)
	for x, line := range field {
		for y, val := range line {
			leftx := x-1
			rightx := x+1
			upy := y-1
			downy := y+1

			// if (x1,y0) exists and f(x1,y0) <= f(x0,y0), then (x0,y0) is **not** a local min
			// if (x0,y1)... ditto
			if !( leftx >= 0 && field[leftx][y] <= val ||
	      		rightx <= maxx && field[rightx][y] <= val ||
	     	  	upy >= 0 && field[x][upy] <= val ||
	      		downy <= maxy && field[x][downy] <= val ) {
	      explored[[2]int{x, y}] = true
	      basin_sizes[[2]int{x, y}] = 1
				risk_level += 1 + val
			}
		}
	}

	// now from each local minimum, explore outward and determine the size of the local basin
	for init, _ := range basin_sizes {
		explore(init, init, maxx, maxy, field, basin_sizes, explored)
	}

	// find the 3 largest basins
	var biggest []int
	for _, val := range basin_sizes {
		biggest = append(biggest, val)
	}
	sort.Ints(biggest)

	fmt.Println("PART 1:", risk_level)
	fmt.Println("PART 2:", biggest[len(biggest)-1] * biggest[len(biggest)-2] * biggest[len(biggest)-3])
}
