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

// read file and return minx, maxx, miny, maxy as ints
func readInput(path string) (int, int, int, int) {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	if scanner.Err() != nil {
		log.Fatalf("scanner: %s", scanner.Err())
	}

	re := regexp.MustCompile(`-?\d+`)
	strnums := re.FindAllString(line, -1)

	nums := make([]int, 4)
	for i, num := range strnums {
		nums[i], _ = strconv.Atoi(num)
	}

	return nums[0], nums[1], nums[2], nums[3]
}

func main() {
	minx, maxx, miny, maxy := readInput("input.txt")

	// track highest y and number of hits on target
	topy, hits := 0, 0

	// since minx > 0 there's no reason to try xvel < 0
	// and xvel > maxx will always overshoot
	for initxvel := 0; initxvel <= maxx; initxvel++ {

		// somewhat arbitrarily chose to try yvel between +/-1000
		yvel_loop:
		for inityvel := -1000; inityvel <= 1000; inityvel++ {

			// track position and highest y per iteration
			x, y, itopy := 0, 0, 0
			xvel, yvel := initxvel, inityvel

			// track position and changing velocity over time
			for true {
				x += xvel
				y += yvel
				if y > itopy { itopy = y }
				if xvel > 0 { xvel-- }
				yvel--

				// overshot top right corner, increasing yvel will only overshoot more
				// OR undershot left edge, increasing yvel will not help
				if (x > maxx && y >= maxy) || (xvel == 0 && x < minx) {
					break yvel_loop

				// overshot x OR overshot y
				} else if x > maxx || y < miny {
					break

				// within the bounds
				} else if x >= minx && x <= maxx && y >= miny && y <= maxy {
					hits++

					// tracking peak y value
					if itopy > topy {
						topy = itopy
					}
					break
				}
			}
		}
	}

	fmt.Println("PART 1:", topy)
	fmt.Println("PART 2:", hits)
}
