package main

import (
	"strings"
	"testing"

	"github.com/josh-weston/advent-of-code-2022/day3/part1"
)

func TestPart1(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`
	stringReader := strings.NewReader(input)
	total := part1.Run(stringReader)

	if total != 157 {
		t.Fatalf("Incorrect value received, want=%d,got=%d", 157, total)
	}

}
