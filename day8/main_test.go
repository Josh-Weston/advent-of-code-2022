package main

import (
	"strings"
	"testing"

	"github.com/josh-weston/advent-of-code-2022/day8/part1"
	"github.com/josh-weston/advent-of-code-2022/day8/part2"
)

func TestConversion(t *testing.T) {

	input := "30373"
	converted := part1.ConvertToIntSlice(input)
	expected := []int{3, 0, 3, 7, 3}

	for i, len := 0, len(converted); i < len; i++ {
		if converted[i] != expected[i] {
			t.Fatalf("Slices are not equal. got=%v, want=%v", converted, expected)
		}
	}
}

func TestHiddenInRowOrColumn(t *testing.T) {
	tests := []struct {
		input    []int
		index    int
		expected bool
	}{
		{[]int{2, 5, 5, 1, 2}, 1, false},
		{[]int{2, 5, 5, 1, 2}, 2, false},
		{[]int{2, 5, 5, 1, 2}, 3, true},
		{[]int{3, 5, 3, 5, 3}, 1, false},
		{[]int{3, 5, 3, 5, 3}, 2, true},
		{[]int{3, 5, 3, 5, 3}, 3, false},
	}

	for _, tt := range tests {
		result := part1.HiddenInRowOrColumn(tt.index, tt.input)
		if result != tt.expected {
			t.Fatalf("Invalid result. got=%t, want=%t", result, tt.expected)
		}
	}

}

func TestPart1(t *testing.T) {
	input := `30373
25512
65332
33549
35390`

	reader := strings.NewReader(input)
	result := part1.Run(reader)
	expected := 21

	if result != expected {
		t.Fatalf("Invalid result. got=%d,want=%d\n", result, expected)
	}

}

func TestPart2(t *testing.T) {
	input := `30373
25512
65332
33549
35390`

	reader := strings.NewReader(input)
	result := part2.Run(reader)
	expected := 8

	if result != expected {
		t.Fatalf("Invalid result. got=%d,want=%d\n", result, expected)
	}

}
