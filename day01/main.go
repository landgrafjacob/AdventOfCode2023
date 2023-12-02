package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file")
	}

	scanner := bufio.NewScanner(f)

	answer1 := 0
	answer2 := 0

	for scanner.Scan() {
		lineValuePart1 := getCalValue(scanner.Text(), false)
		lineValuePart2 := getCalValue(scanner.Text(), true)
		answer1 = answer1 + lineValuePart1
		answer2 = answer2 + lineValuePart2
		// fmt.Printf("%s: %d\n", scanner.Text(), lineValue)
	}

	fmt.Printf("Calibration value Part 1: %d\n", answer1)
	fmt.Printf("Calibration value Part 2: %d\n", answer2)
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
