package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

// take a single of the format = "1 2 3 4 5"
// return an []int = [1 2 3 4 5]
func stringToInts(line string) ([]int) {
	re := regexp.MustCompile(`[, ]`)
	string_array := re.Split(line, -1)
	var int_array []int

	for _, number := range string_array {
		if number == "" {
			continue
		}
		val, err := strconv.Atoi(number)
		if err != nil {
			log.Fatalf("strconv.Atoi: %s", err)
		}

		int_array = append(int_array, val)
	}

	return int_array
}

// take raw format of lines and interpret it into:
// draws ([]int) the ordered bingo draws, and
// cards ([i][j][k]int) the bingo cards, where
// i is a bingo card number
// j, k are the rows, columns of the numbers on a bingo card
// eg. cards[0] = {{1 2 3 4 5} {6 7 8 9 10} {11 12 13 14 15} ...}
func processLines(lines []string) (draws []int, cards [][][]int) {
	// first line = draws
	draws = stringToInts(lines[0])

	// second line = empty, disregard
	// lines 3-7 = cards[0]
	// line 8 = empty
	// lines 9-13 = cards[1]
	// ...
	lines = lines[2:]
	i := 0

	for j := 0; j < len(lines); j += 6 {
		var card [][]int
		for _, line := range lines[j:j+5] {
			card = append(card, stringToInts(line))
		}
		cards = append(cards, card)
		i++
	}

	return draws, cards
}

// given a drawn number, mark a card with -1 in place of the original number
func markCard(draw int, card [][]int) ([][]int) {
	for i, row := range card {
		for j, num := range row {
			if num == draw {
				card[i][j] = -1
				return card
			}
		}
	}
	return card
}

// if all values in a line are -1, we have a winner!
func checkLineWinner(line []int) (bool) {
	for _, num := range line {
		if num != -1 {
			return false
		}
	}

	return true
}

// take an n x m array and return its m x n transpose
func transpose(card [][]int) ([][]int) {
	var transposed [][]int

	for i, row := range card {
		var transposed_row []int
		for j, _ := range row {
			transposed_row = append(transposed_row, card[j][i])
		}
		transposed = append(transposed, transposed_row)
	}

	return transposed
}

// check each row and column for possible wins
func checkCardWinner(card [][]int) (bool) {
	for _, row := range card {
		if checkLineWinner(row) {
			return true
		}
	}

	// rather than creating arrays for each column, I lazily transpose and check rows again
	card = transpose(card)
	for _, row := range card {
		if checkLineWinner(row) {
			return true
		}
	}

	return false
}

// sum any unmarked numbers and multiply by the winning draw
func calculateScore(draw int, card [][]int) (int) {
	sum := 0

	for _, row := range card {
		for _, num := range row {
			if num != -1 {
				sum += num
			}
		}
	}

	return sum * draw
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	draws, cards := processLines(lines)

	// I need to preserve these variables after the loop has broken so they're declared here
	var draw int
	var i int
	var j int
	var card [][]int

	// find the first winner
	loop_one:
	for i, draw = range draws {
		for j, _ = range cards {
			card = markCard(draw, cards[j])

			// here they are!
			if checkCardWinner(card) {
				break loop_one
			}
		}
	}

	fmt.Println("PART 1:")
	fmt.Println(calculateScore(draw, card))

	// remove the winning card from the stack
	cards = append(cards[:j], cards[j + 1:]...)

	// now continue until all cards have won
	// I'm replaying the entire first winning round, as opposed to continuing where I left off, but so what?
	loop_two:
	for _, draw = range draws[i:] {

		// range cards won't account for changes to cards mid-loop, so track the number of removed cards
		removed := 0

		for j, _ = range cards {
			card = markCard(draw, cards[j - removed])
			if checkCardWinner(card) {

				// the final winner!
				if len(cards) == 1 {
					break loop_two
				}

				cards = append(cards[:j - removed], cards[j - removed + 1:]...)
				removed++
			}
		}
	}

	fmt.Println("PART 2:")
	fmt.Println(calculateScore(draw, card))
}
