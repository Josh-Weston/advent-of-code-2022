package main

import (
	"strings"
	"testing"

	"github.com/josh-weston/advent-of-code-2022/day12/part1"
	"github.com/josh-weston/advent-of-code-2022/day12/part2"
)

// a is lowest elevation, z is heighest elevation
// S = current position (a)
// E = best position (z)

// reach E in as few steps as possible
// can move N/S and E/W, can only move to an elevation 1 higher than your current, or any elevation <= your current elevation

// 1. Randomize the moves and run it a bunch of times (one option)
// 2. Use a genetic algorithm to randomize and evolve

func TestPart1(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

	reader := strings.NewReader(input)
	result := part1.BFS(reader)
	expected := 31
	if result != 31 {
		t.Fatalf("Invalid result. got=%d,want=%d", result, expected)
	}

}

func TestPart2(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

	reader := strings.NewReader(input)
	result := part2.Run(reader)
	expected := 29
	if result != 29 {
		t.Fatalf("Invalid result. got=%d,want=%d", result, expected)
	}

}
