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
	"strings"
)

// read file and process lines into a map[string][]string
// where keys are positions on the map and values are all positions directly connected
func readInput(path string) (map[string][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cave_map := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		cave_map[line[0]] = append(cave_map[line[0]], line[1])
		cave_map[line[1]] = append(cave_map[line[1]], line[0])
	}
	return cave_map, scanner.Err()
}

// Given a map of the caves, the current position, and a list of small caves already visited,
// proceed in all valid directions and return the number of such paths which results in `end`
func findPathsPartOne(cave_map map[string][]string, position string, visited map[string]bool) (paths int) {

	// terminates a path
	if position == "end" {
		return 1
	}

	// rassafrassin golang forcing me to deep copy...
	visited_copy := make(map[string]bool)
	for k, v := range visited {
		visited_copy[k] = v
	}

	// count up the branches that result in "end"
	sum := 0
	visited_copy[position] = true
	for _, direction := range cave_map[position] {

		// lowercase, already visited? can't go this way
		if direction[0] >= 97 && visited[direction] {
			continue
		}
		sum += findPathsPartOne(cave_map, direction, visited_copy)
	}

	return sum
}

// Given a map of the caves, the current position, counts of the number of times each cave has been visited,
// and a bool
// proceed in all valid directions and return the number of such paths which results in `end`
func findPathsPartTwo(cave_map map[string][]string, position string, visited map[string]int, been_twice bool) (paths int) {
	if position == "end" {
		return 1
	}

	visited_copy := make(map[string]int)
	for k, v := range visited {
		visited_copy[k] = v
	}

	sum := 0
	visited_copy[position]++

	// check whether we've paid a second visit to a small cave
	been_twice = been_twice || (position != "start" && position[0] >= 97 && visited_copy[position] >= 2)
	for _, direction := range cave_map[position] {

		// lowercase
		if direction[0] >= 97 {

			// cannot return to "start" or visit a small cave a second time if we've previously visited one twice
			if direction == "start" || (been_twice && visited_copy[direction] >= 1) {
				continue
			} else {
				sum += findPathsPartTwo(cave_map, direction, visited_copy, been_twice)
			}

		// uppercase
		} else {
			sum += findPathsPartTwo(cave_map, direction, visited_copy, been_twice)
		}
	}

	return sum
}

func main() {
	cave_map, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readInput: %s", err)
	}

	visit_once := make(map[string]bool)
	visit_twice := make(map[string]int)
	fmt.Println("PART 1:", findPathsPartOne(cave_map, "start", visit_once))
	fmt.Println("PART 2:", findPathsPartTwo(cave_map, "start", visit_twice, false))
}
