package part2

import "github.com/josh-weston/advent-of-code-2022/day11/types"

// how worried you are about each item

// monkeys have attributes (classes)
// worry level for each item in the order they will be inspected
// after inspecting, divide by 3 and round down to the nearest integer
// if a monkey has no items, it will inspect them as they come through

// round = all monkeys taking a single turn
// total number of times each monkey inspects items over 20 rounds
// multiply their numbers together

func Run(orchestrator types.MonkeyOrchestrator2, rounds int) int {
	orchestrator.Begin(rounds)
	return orchestrator.Score()
}
