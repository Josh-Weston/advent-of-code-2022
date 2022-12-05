package main

import (
	"strings"
	"testing"

	"github.com/josh-weston/advent-of-code-2022/day5/part1"
	"github.com/josh-weston/advent-of-code-2022/day5/part2"
)

// the crates need to be rearranged
// crange to rearrange crates so the desired crates are at the top
// the crates move one at a time, but they might move multiple crates in a single instruction
// the instructions are 1 indexed
func TestPart1(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	reader := strings.NewReader(input)
	solution := part1.Run(reader, 3)
	expected := "CMZ"

	if solution != expected {
		t.Fatalf("incorrect solution. want=%s, got=%s", expected, solution)
	}

}

func TestPart2(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	reader := strings.NewReader(input)
	solution := part2.Run(reader, 3)
	expected := "MCD"

	if solution != expected {
		t.Fatalf("incorrect solution. want=%s, got=%s", expected, solution)
	}

}
