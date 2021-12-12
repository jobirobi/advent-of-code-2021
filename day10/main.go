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

// read file and return lines as []string
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

func reverse(chars []rune) {
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	syntax_score := 0
	var autocorrect_scores []int
	for _, line := range lines {
		var stack []rune
		incorrupt := true
		char_loop:
		for _, char := range line {
			switch char {
			/*
				ascii values:
				40 = (
				41 = )
				60 = <
				62 = >
				91 = [
				93 = ]
				123 = {
				125 =	}

				if char is an opening character, add to the stack
				if char is a closing character and is the mate of the top of the stack, remove mate from the stack
				if char is a closing character and is not the mate, line is corrupted
			*/
			case 40, 60, 91, 123:
				stack = append(stack, char)
			case 41:
				if stack[len(stack)-1] != 40 {
					syntax_score += 3
					incorrupt = false
					break char_loop
				} else {
					stack = stack[:len(stack)-1]
				}
			case 62:
				if stack[len(stack)-1] != 60 {
					syntax_score += 25137
					incorrupt = false
					break char_loop
				} else {
					stack = stack[:len(stack)-1]
				}
			case 93:
				if stack[len(stack)-1] != 91 {
					syntax_score += 57
					incorrupt = false
					break char_loop
				} else {
					stack = stack[:len(stack)-1]
				}
			case 125:
				if stack[len(stack)-1] != 123 {
					syntax_score += 1197
					incorrupt = false
					break char_loop
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
		if incorrupt {
			score := 0
			reverse(stack)
			for _, char := range stack {
				switch char {
				case 40:
					score = 5 * score + 1
				case 60:
					score = 5 * score + 4
				case 91:
					score = 5 * score + 2
				case 123:
					score = 5 * score + 3
				}
			}
			autocorrect_scores = append(autocorrect_scores, score)
		}
	}
	sort.Ints(autocorrect_scores)

	fmt.Println("PART 1:", syntax_score)
	fmt.Println("PART 2:", autocorrect_scores[len(autocorrect_scores) / 2])
}
