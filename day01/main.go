package main

import (
	"fmt"
	"regexp"
	"strconv"
	"bufio"
	"flag"

	"github.com/landgrafjacob/AdventOfCode2023/utils"
)

var numMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var part int

func main() {
	scanner := utils.ReadFile("input.txt")
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	if part == 1 {
		fmt.Printf("Calibration value Part 1: %d\n", part01(scanner))
	} else if part == 2 {
		fmt.Printf("Calibration value Part 1: %d\n", part02(scanner))
	}
}

func part01(scanner *bufio.Scanner) int {
	answer := 0
	for scanner.Scan() {
		lineValue := getCalValue(scanner.Text(), false)
		answer += lineValue
	}

	return answer
}

func part02(scanner *bufio.Scanner) int {
	answer := 0
	for scanner.Scan() {
		lineValue := getCalValue(scanner.Text(), true)
		answer += lineValue
	}

	return answer
}

func getFirstNum(line string, includeSpelled bool) string {
	for i, char := range line {
		charString := string(char)
		if regexp.MustCompile(`\d`).MatchString(charString) {
			return charString
		}
		if includeSpelled {
			for j := 3; j <= 5; j++ {
				if i+j > len(line) {
					continue
				}
				if num, ok := numMap[line[i:i+j]]; ok {
					return num
				}
			}
		}
	}
	return ""
}

func getLastNum(line string, includeSpelled bool) string {
	for i := len(line) - 1; i >= 0; i-- {
		charString := string(line[i])
		if regexp.MustCompile(`\d`).MatchString(charString) {
			return charString
		}
		if includeSpelled {
			for j := 3; j <= 5; j++ {
				if i+j > len(line) {
					continue
				}
				if num, ok := numMap[line[i:i+j]]; ok {
					return num
				}
			}
		}
	}
	return ""
}

func getCalValue(line string, includeSpelled bool) int {
	first := getFirstNum(line, includeSpelled)
	second := getLastNum(line, includeSpelled)

	result, _ := strconv.Atoi(first + second)
	return result
}
