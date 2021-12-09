// I'll just come out and say it - I did things pretty sloppily today. I still solved it, so...

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	// "strconv"
	"strings"
)

// TIL you gotta write your own string sort in golang
type stringSort []rune
func (s stringSort) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s stringSort) Less(i, j int) bool { return s[i] < s[j] }
func (s stringSort) Len() int { return len(s) }

// take a string with "abcdf" in it and determine what "eg" is
func findEG(abcdf string) (string) {
	eg := []rune("abcdefg")
	del := []rune(abcdf)
	sort.Sort(stringSort(del))
	t := 0
	for _, c := range del {
		if strings.ContainsRune(string(eg), c) {
			switch c {
			case []rune("a")[0]:
				eg = eg[1:]
				t++
			case []rune("b")[0]:
				eg = append(eg[:1-t], eg[1-t+1:]...)
				t++
			case []rune("c")[0]:
				eg = append(eg[:2-t], eg[2-t+1:]...)
				t++
			case []rune("d")[0]:
				eg = append(eg[:3-t], eg[3-t+1:]...)
				t++
			case []rune("e")[0]:
				eg = append(eg[:4-t], eg[4-t+1:]...)
				t++
			case []rune("f")[0]:
				eg = append(eg[:5-t], eg[5-t+1:]...)
				t++
			case []rune("g")[0]:
				eg = eg[:len(eg)-1]
			}
		}
	}

	return string(eg)
}

/*
	take strings for "cf" and "eg" to determine what a string maps to

	length 5, containing both "eg" = 2
	length 5, containing only "g" and
		containing both "cf" = 3
		containing only "f"  = 5
	length 6, containing only "g"  = 9
	length 6, containing both "eg" and
		containing both "cf" = 0
		containing only "f"  = 6
*/
func findN(cf string, eg string, id_me string) (int) {
	switch len(id_me) {
	case 5:
		if strings.ContainsRune(id_me, []rune(eg)[0]) && strings.ContainsRune(id_me, []rune(eg)[1]) {
			return 2
		} else if strings.ContainsRune(id_me, []rune(cf)[0]) && strings.ContainsRune(id_me, []rune(cf)[1]) {
			return 3
		} else {
			return 5
		}
	case 6:
		if !(strings.ContainsRune(id_me, []rune(eg)[0]) && strings.ContainsRune(id_me, []rune(eg)[1])) {
			return 9
		} else if strings.ContainsRune(id_me, []rune(cf)[0]) && strings.ContainsRune(id_me, []rune(cf)[1]) {
			return 0
		} else {
			return 6
		}
	default:
		fmt.Println("Something went wrong")
		return 10
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("readInput: %s", err)
	}
	defer file.Close()

	total_pt1, total_pt2 := 0,0

	re := regexp.MustCompile(`[ |]+`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var cf, acf, bcdf string
		segments := make(map[string]int)

		/*
			DIAGRAM:
		   aaaa
		  b    c
		  b    c
		   dddd
		  e    f
		  e    f
		   gggg

		  SEGMENTS MAP:
		  "abcefg"  = 0
		  "cf"      = 1
		  "acdeg"   = 2
		  "acdfg"   = 3
		  "bcdf"    = 4
		  "abdfg"   = 5
		  "abdefg"  = 6
		  "acf"     = 7
		  "abcdefg" = 8
		  "abcdfg"  = 9
		*/

		line := re.Split(scanner.Text(), -1)
		var unidentified []string
		for _, str := range line[0:10] {

			// sorting the strings because sometimes multiple permutations appear in the same line, eg. ab ba
			s := []rune(str)
			sort.Sort(stringSort(s))

			switch len(str) {
			case 2:
				segments[string(s)] = 1
				cf = string(s)
			case 3:
				segments[string(s)] = 7
				acf = string(s)
			case 4:
				segments[string(s)] = 4
				bcdf = string(s)
			case 7:
				segments[string(s)] = 8
			default:
				unidentified = append(unidentified, string(s))
			}
		}

		/*
		  at this point we can identify:
		  the chars corresponding to "abcdf" in the diagram,
		  leaving "eg" unidentified.

		  step 1: determine which two chars correspond to "eg"
		  step 2: use known "cf" and "eg" to identify other numbers
		*/
		eg := findEG(acf + bcdf)
		for _, id_me := range unidentified {
			segments[id_me] = findN(cf, eg, id_me)
		}

		// now count up the number of occurrences 1,4,7,8 in the right side (part 1)
		// and the sum the values of the right side (part 2)
		for i, str := range line[10:] {
			s := []rune(str)
			sort.Sort(stringSort(s))

			switch len(str) {
			case 2:
				total_pt1++
			case 3:
				total_pt1++
			case 4:
				total_pt1++
			case 7:
				total_pt1++
			}
			total_pt2 += int(math.Pow(float64(10), float64(3 - i)) * float64(segments[string(s)]))
		}
	}

	fmt.Println("PART 1:", total_pt1)
	fmt.Println("PART 2:", total_pt2)
}
