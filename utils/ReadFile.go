package utils

import (
	"os"
	"fmt"
	"bufio"
)

func ReadFile (fileName string) *bufio.Scanner {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening input file")
	}

	return bufio.NewScanner(f)
}