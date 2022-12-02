package main

import (
	"strings"
	"testing"

	"github.com/josh-weston/advent-of-code-2022/day2/part1"
)

func TestInput(t *testing.T) {
	input := `A Y
B X
C Z
`
	reader := strings.NewReader(input)
	score := part1.Run(reader)
	if score != 15 {
		t.Fatalf("Incorrect score. got=%d, want=%d", score, 15)
	}

}
