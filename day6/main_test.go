package main

import (
	"strings"
	"testing"

	"github.com/josh-weston/advent-of-code-2022/day6/part1"
	"github.com/josh-weston/advent-of-code-2022/day6/part2"
)

func TestPart1(t *testing.T) {

	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 5,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 6,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 10,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 11,
		},
		{
			input:    "bbbbbabcd",
			expected: 9,
		},
		{
			input:    "abcdeeeeeeee",
			expected: 4,
		},
	}

	for _, tt := range tests {
		input := strings.NewReader(tt.input)
		result := part1.Run(input)
		if result != tt.expected {
			t.Fatalf("unexpected result. got=%d, want=%d", result, tt.expected)
		}
	}

}

func TestPart2(t *testing.T) {

	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 19,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 23,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 23,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 29,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 26,
		},
	}

	for _, tt := range tests {
		input := strings.NewReader(tt.input)
		result := part2.Run(input)
		if result != tt.expected {
			t.Fatalf("unexpected result. got=%d, want=%d", result, tt.expected)
		}
	}

}
