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

func readInput(path string) (string, map[string]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	var chain string
	rules := make(map[string]rune)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// skip empty lines
		if line == "" {
			continue
		}
		splits := strings.Split(line, " -> ")

		// rule lines, eg "AB -> C"
		if len(splits) == 2 {
			rules[splits[0]] = []rune(splits[1])[0]

		// starting chain line, eg "ABCD"
		} else {
			chain = line
		}
	}
	return chain, rules, scanner.Err()
}

func main() {
	chain, rules, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readInput: %s", err)
	}
	first, last := chain[0], chain[len(chain)-1]

	// interpret chain into character-pairs and quantities
	pairs := make(map[string]int)
	for i := 0; i < len(chain) - 1; i++ {
		pairs[chain[i:i+2]]++
	}

	// build the chain without caring about order
	i := 1
	for ; i <= 10; i++ {
		deepcopy := make(map[string]int)
		for key, val := range pairs {
			deepcopy[key] = val
		}
		for key, val := range pairs {

			// chain = "AB" => "ACB", pairs = {"AB":1} => {"AB":0, "AC":1, "CB":1}
			deepcopy[key] -= val
			deepcopy[string(key[0]) + string(rules[key])] += val
			deepcopy[string(rules[key]) + string(key[1])] += val
		}
		pairs = deepcopy
	}

	singles := map[byte]int{
		first: 1,
		last: 1,
	}
	for key, val := range pairs {
		singles[[]byte(key)[0]] += val
		singles[[]byte(key)[1]] += val
	}
	// at this point each single has been counted exactly double
	// eg. chain = "ABCD", first = "A", last = "D" pairs = {"AB":1, "BC":1, "CD":1}
	// singles = {"A":2, "B":2, "C": 2, "D":2}
	least, most := 9223372036854775807, 0
	for _, val := range singles {
		if least > val / 2 {
			least = val / 2
		}
		if most < val / 2 {
			most = val / 2
		}
	}
	fmt.Println("PART 1:", most - least)

	// keep going to 40th loop
	for ; i <= 40; i++ {
		deepcopy := make(map[string]int)
		for key, val := range pairs {
			deepcopy[key] = val
		}
		for key, val := range pairs {
			deepcopy[key] -= val
			deepcopy[string(key[0]) + string(rules[key])] += val
			deepcopy[string(rules[key]) + string(key[1])] += val
		}
		pairs = deepcopy
	}
	singles = map[byte]int{
		first: 1,
		last: 1,
	}
	for key, val := range pairs {
		singles[[]byte(key)[0]] += val
		singles[[]byte(key)[1]] += val
	}
	// at this point each single has been counted exactly double
	// eg. chain = "ABCD", first = "A", last = "D" pairs = {"AB":1, "BC":1, "CD":1}
	// singles = {"A":2, "B":2, "C": 2, "D":2}
	least, most = 9223372036854775807, 0
	for _, val := range singles {
		if least > val / 2 {
			least = val / 2
		}
		if most < val / 2 {
			most = val / 2
		}
	}
	fmt.Println("PART 2:", most - least)
}
