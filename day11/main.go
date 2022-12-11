package main

import (
	"fmt"

	// "github.com/josh-weston/advent-of-code-2022/day11/part1"
	"github.com/josh-weston/advent-of-code-2022/day11/part2"
	"github.com/josh-weston/advent-of-code-2022/day11/types"
)

func main() {
	var orchestrator types.MonkeyOrchestrator2C
	orchestrator.Items = [][]int{
		{98, 89, 52},
		{57, 95, 80, 92, 57, 78},
		{82, 74, 97, 75, 51, 92, 83},
		{97, 88, 51, 68, 76},
		{63},
		{94, 91, 51, 63},
		{61, 54, 94, 71, 74, 68, 98, 83},
		{90, 56},
	}

	orchestrator.Inspected = make([]int, len(orchestrator.Items))

	orchestrator.Operations = []func(i int) int{
		func(v int) int {
			return v * 2
		},
		func(v int) int {
			return v * 13
		},
		func(v int) int {
			return v + 5
		},
		func(v int) int {
			return v + 6
		},
		func(v int) int {
			return v + 1
		},
		func(v int) int {
			return v + 4
		},
		func(v int) int {
			return v + 2
		},
		func(v int) int {
			return v * v
		},
	}

	orchestrator.Tests = []func(i int) int{
		func(v int) int {
			if v%5 == 0 {
				return 6
			}
			return 1
		},
		func(v int) int {
			if v%2 == 0 {
				return 2
			}
			return 6
		},
		func(v int) int {
			if v%19 == 0 {
				return 7
			}
			return 5
		},
		func(v int) int {
			if v%7 == 0 {
				return 0
			}
			return 4
		},
		func(v int) int {
			if v%17 == 0 {
				return 0
			}
			return 1
		},
		func(v int) int {
			if v%13 == 0 {
				return 4
			}
			return 3
		},
		func(v int) int {
			if v%3 == 0 {
				return 2
			}
			return 7
		},
		func(v int) int {
			if v%11 == 0 {
				return 3
			}
			return 5
		},
	}

	result := part2.Run2C(orchestrator, 10_000)
	fmt.Printf("The result is %d\n", result)
}
