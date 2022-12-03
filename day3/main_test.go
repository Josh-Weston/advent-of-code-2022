package main

import (
	"strings"
	"testing"

	"github.com/josh-weston/advent-of-code-2022/day3/part1"
	"github.com/josh-weston/advent-of-code-2022/day3/part2"
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

func TestPart2(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`
	stringReader := strings.NewReader(input)
	total := part2.Run(stringReader)
	expected := 18
	if total != expected {
		t.Fatalf("Incorrect value received, want=%d,got=%d", expected, total)
	}

}
