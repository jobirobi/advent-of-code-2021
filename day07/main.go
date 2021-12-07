package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

// arithmetic mean
func mean(nums []int) (float64) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return float64(total) / float64(len(nums))
}

// arithmetic median
func median(nums []int) (float64) {
	sort.Ints(nums)
	if len(nums) % 2 == 1 {
		return float64(nums[len(nums) / 2])
	} else {
		return mean(nums[len(nums) / 2 - 1:len(nums) / 2])
	}
}

// fuel required in the part 1 scenario
// I got lucky when I happened to notice that the optimal point in the test case == the median of the set
// so I tried extrapolating, and it worked!
func fuelRequiredConstant(nums []int, point float64) (int) {
	var total float64
	for _, num := range nums {
		total += math.Abs(point - float64(num))
	}
	return int(total)
}

// fuel required in the part 2 scenario
// I got lucky when I happened to notice that the optimal point in the test case == the mean of the set
// so again I tried extrapolating, and it worked!
func fuelRequiredVariable(nums []int, point float64) (int) {
	var total float64
	for _, num := range nums {
		n := math.Abs(point - float64(num))

		// Faulhaber’s formula (k=1->n)Σ(k) = n(n+1)/2
		total += n * (n + 1.0) / 2.0
	}
	return int(total)
}

func main() {
	nums, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("readInput: %s", err)
	}

	// calculate the median, and round it both up and down
	fmt.Println("PART 1")
	avg := median(nums)
	low, high := math.Floor(avg), math.Ceil(avg)

	// now determine which of the round-up and round-down cases is better
	if low == high {
		fmt.Println("The optimal point is", low)
		fmt.Println("Minimum fuel required:", fuelRequiredConstant(nums, low))
	} else {
		low_fuel, high_fuel := fuelRequiredConstant(nums, low), fuelRequiredConstant(nums, high)
		if low_fuel <= high_fuel {
			fmt.Println("The optimal point is", low)
			fmt.Println("Minimum fuel required:", low_fuel)
		} else {
			fmt.Println("The optimal point is", high)
			fmt.Println("Minimum fuel required:", high_fuel)
		}
	}

	// calculate the mean, and round it both up and down
	fmt.Println("PART 2")
	avg = mean(nums)
	low, high = math.Floor(avg), math.Ceil(avg)

	// now determine which of the round-up and round-down cases is better
	if low == high {
		fmt.Println("The optimal point is", low)
		fmt.Println("Minimum fuel required:", fuelRequiredVariable(nums, low))
	} else {
		low_fuel, high_fuel := fuelRequiredVariable(nums, low), fuelRequiredVariable(nums, high)
		if low_fuel <= high_fuel {
			fmt.Println("The optimal point is", low)
			fmt.Println("Minimum fuel required:", low_fuel)
		} else {
			fmt.Println("The optimal point is", high)
			fmt.Println("Minimum fuel required:", high_fuel)
		}
	}
}
