package main

import (
	"testing"

	"github.com/landgrafjacob/AdventOfCode2023/utils"
)

func Test_part01(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected int
	}{
		{
			name: "test",
			input: "input_test.txt",
			expected: 142,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			scanner := utils.ReadFile(test.input)
			result := part01(scanner)
			if result != test.expected {
				t.Errorf("Part1 Error - Expected: %v, Got: %v", test.expected, result)
			}
		})
	}
}

func Test_part02(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected int
	}{
		{
			name: "test",
			input: "input_test_2.txt",
			expected: 281,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			scanner := utils.ReadFile(test.input)
			result := part02(scanner)
			if result != test.expected {
				t.Errorf("Part1 Error - Expected: %v, Got: %v", test.expected, result)
			}
		})
	}
}