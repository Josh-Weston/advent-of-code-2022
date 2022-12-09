package main

import (
	"strings"
	"testing"

	"github.com/josh-weston/advent-of-code-2022/day9/part1"
)

func TestPart1(t *testing.T) {

	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	reader := strings.NewReader(input)
	result := part1.Run(reader)
	expected := 13

	if result != expected {
		t.Fatalf("invalid result, want=%d,got=%d", expected, result)
	}

}
