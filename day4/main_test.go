package main

import (
	"strings"
	"testing"

	"github.com/josh-weston/advent-of-code-2022/day4/part1"
	"github.com/josh-weston/advent-of-code-2022/day4/part2"
)

func TestPart1(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

	reader := strings.NewReader(input)
	expected := 2
	n := part1.Run(reader)
	if n != expected {
		t.Fatalf("Incorrect value received, want=%d,got=%d", expected, n)
	}
}

func TestPart2(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

	reader := strings.NewReader(input)
	expected := 4
	n := part2.Run(reader)
	if n != expected {
		t.Fatalf("Incorrect value received, want=%d,got=%d", expected, n)
	}
}

func TestPartManual(t *testing.T) {
	input := `28-47,45-47
32-97,98-98
59-92,91-93
8-74,9-74
58-98,23-57
`

	reader := strings.NewReader(input)
	expected := 3
	n := part2.Run(reader)
	if n != expected {
		t.Fatalf("Incorrect value received, want=%d,got=%d", expected, n)
	}
}
