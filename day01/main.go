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
	f, err := os.Open("input_test_2.txt")
	if err != nil {
		fmt.Println("Error opening input file")
	}

	scanner := bufio.NewScanner(f)

	answer := 0

	for scanner.Scan() {
		lineValue := getCalValue(scanner.Text())
		answer = answer + lineValue
		fmt.Printf("%s: %d\n", scanner.Text(), lineValue)
	}

	fmt.Printf("Calibration value: %d\n", answer)
}

func getFirstNum(line string) string {
	for i, char := range line {
		charString := string(char)
		if regexp.MustCompile(`\d`).MatchString(charString) {
			return charString
		}
		for j := 3; j <= 5; j++ {
			if num, ok := numMap[line[i:i+3]]; ok {
				return num
			}
		}
	}
	return ""
}

func getLastNum(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		charString := string(line[i])
		if regexp.MustCompile(`\d`).MatchString(charString) {
			return charString
		}
	}
	return ""
}

func getCalValue(line string) int {
	first := getFirstNum(line)
	second := getLastNum(line)
	fmt.Printf("%s: %s", line, first)

	result, _ := strconv.Atoi(first + second)
	return result
}
