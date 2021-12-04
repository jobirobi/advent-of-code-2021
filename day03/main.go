package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

// this function doubles as a [1]string to []int converter!
func countOnes(array []string) ([]int) {
	// to initialize an array of not-a-constant length, use make()
	ones_counts := make([]int, len(array[0]))

	for _, line := range array {
		for i, char := range line {
			if string(char) == "1" {
				ones_counts[i]++
			}
		}
	}

	return ones_counts
}

// determine gamma value and epsilon value of an array of numbers
// gamma = more common digit per index
// epsilon = less common digit per index
// eg. [0101, 0101, 1001, 0001, 1110] => gamma = 0101, epsilon = 1010
func gammaEpsilon(lines []string) ([]int, []int) {
	ones_counts := countOnes(lines)
	num_lines := len(lines)

	// initializes with all 0s, eg. 0000
	gamma_binary := make([]int, len(ones_counts))
	epsilon_binary := make([]int, len(ones_counts))

	for i, count := range ones_counts {
		// in a tie between 0 and 1, gamma = 1, epsilon = 0
		if float32(count) >= (float32(num_lines) / 2.0) {
			gamma_binary[i] = 1
		} else {
			epsilon_binary[i] = 1
		}
	}

	return gamma_binary, epsilon_binary
}

// binary to decimal function
// least significant digit = 2^0, second-least = 2^1, third-least = 2^2, etc.
// take binary array, eg. 1010, and add the value of each digit to calculate decimal
// eg. 1 * 2^3 + 0 * 2^2 + 1 * 2^1 + 0 * 2^0 = 10
func binaryArrayToDecimal(array []int) (int) {
	decimal := 0

	for i, digit := range array {
		decimal += digit * int(math.Pow(float64(2), float64(len(array) - 1 - i)))
	}

	return decimal
}

// take []string and reduce iteratively to get [1]string O2 generator rating, [1]string CO2 scrubber rating
// ... according to the rules in the problem statement, not going to repeat them here
func reduceLines(gamma_lines []string) ([]string, []string) {
	epsilon_lines := gamma_lines

	// first calculate the more/less common digits per index
	gamma_binary, epsilon_binary := gammaEpsilon(gamma_lines)

	// going one digit at a time
	for i, _ := range gamma_binary {

		// determine the filter keys
		gamma_keeper := strconv.Itoa(gamma_binary[i])
		epsilon_keeper := strconv.Itoa(epsilon_binary[i])
		gamma_kept := []string{}
		epsilon_kept := []string{}

		// now progressing line by line
		for _, line := range gamma_lines {

			// keep a line if digit at index i of line == the keeper value
			if string(line[i]) == gamma_keeper {
				gamma_kept = append(gamma_kept, line)
			}
		}
		for _, line := range epsilon_lines {
			if string(line[i]) == epsilon_keeper {
				epsilon_kept = append(epsilon_kept, line)
			}
		}

		// exit once one line remains in each filter
		if len(gamma_kept) == 1 && len(epsilon_kept) == 1 {
			return gamma_kept, epsilon_kept

		// I had "else if len(gamma_lines) == 1" here but by definition the epsilon filter is more choosy

		} else if len(epsilon_kept) == 1 {
			gamma_lines = gamma_kept
			epsilon_lines = epsilon_kept
			gamma_binary, _ = gammaEpsilon(gamma_lines)

			// maintain epsilon_lines as is
			epsilon_binary = countOnes(epsilon_lines)
		} else {
			gamma_lines = gamma_kept
			epsilon_lines = epsilon_kept
			gamma_binary, _ = gammaEpsilon(gamma_lines)
			_, epsilon_binary = gammaEpsilon(epsilon_lines)
		}
	}

	fmt.Println("Something went wrong")
	return []string{}, []string{}
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	gamma_binary, epsilon_binary := gammaEpsilon(lines)

	// at this point I have []ints for gamma, epsilon representing binary numbers, eg: [0 1 0 1], [1 0 1 0]
	// and I need to convert them into decimal numbers, eg: 5, 10
	gamma_decimal := binaryArrayToDecimal(gamma_binary)
	epsilon_decimal := binaryArrayToDecimal(epsilon_binary)

	fmt.Println("PART ONE:")
	fmt.Println("gamma =", gamma_binary, "=", gamma_decimal)
	fmt.Println("epsilon =", epsilon_binary, "=", epsilon_decimal)
	fmt.Println("gamma * epsilon =", gamma_decimal * epsilon_decimal)

	fmt.Println("")
	fmt.Println("PART TWO:")
	gamma_reduction, epsilon_reduction := reduceLines(lines)
	gamma_decimal = binaryArrayToDecimal(countOnes(gamma_reduction))
	epsilon_decimal = binaryArrayToDecimal(countOnes(epsilon_reduction))

	fmt.Println("O2 rating =", gamma_reduction, "=", gamma_decimal)
	fmt.Println("CO2 rating =", epsilon_reduction, "=", epsilon_decimal)
	fmt.Println("O2 rating * CO2 rating =", gamma_decimal * epsilon_decimal)
}
