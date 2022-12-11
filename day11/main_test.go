package main

import (
	"math/big"
	"testing"

	"github.com/josh-weston/advent-of-code-2022/day11/part1"
	"github.com/josh-weston/advent-of-code-2022/day11/part2"
	"github.com/josh-weston/advent-of-code-2022/day11/types"
)

func TestPart1(t *testing.T) {

	var orchestrator types.MonkeyOrchestrator
	orchestrator.Items = [][]int{
		{79, 98},
		{54, 65, 75, 74},
		{79, 60, 97},
		{74},
	}

	orchestrator.Inspected = []int{0, 0, 0, 0}

	orchestrator.Operations = []func(i int) int{
		func(v int) int {
			return v * 19
		},
		func(v int) int {
			return v + 6
		},
		func(v int) int {
			return v * v
		},
		func(v int) int {
			return v + 3
		},
	}

	orchestrator.Tests = []func(i int) int{
		func(v int) int {
			if v%23 == 0 {
				return 2
			}
			return 3
		},
		func(v int) int {
			if v%19 == 0 {
				return 2
			}
			return 0
		},
		func(v int) int {
			if v%13 == 0 {
				return 1
			}
			return 3
		},
		func(v int) int {
			if v%17 == 0 {
				return 0
			}
			return 1
		},
	}

	result := part1.Run(orchestrator, 20)
	expected := 10605

	if result != expected {
		t.Fatalf("invalid result. got=%d, want=%d\n", result, expected)
	}

}

func TestPart2(t *testing.T) {

	var orchestrator types.MonkeyOrchestrator2
	orchestrator.Items = [][]*big.Int{
		{big.NewInt(79), big.NewInt(98)},
		{big.NewInt(54), big.NewInt(65), big.NewInt(75), big.NewInt(74)},
		{big.NewInt(79), big.NewInt(60), big.NewInt(97)},
		{big.NewInt(74)},
	}

	orchestrator.Inspected = []int{0, 0, 0, 0}

	orchestrator.Operations = []func(v *big.Int) *big.Int{
		func(v *big.Int) *big.Int {
			return v.Mul(v, big.NewInt(19))
		},
		func(v *big.Int) *big.Int {
			return v.Add(v, big.NewInt(6))
		},
		func(v *big.Int) *big.Int {
			return v.Mul(v, v)
		},
		func(v *big.Int) *big.Int {
			return v.Add(v, big.NewInt(3))
		},
	}

	orchestrator.Tests = []func(v *big.Int) int{
		func(v *big.Int) int {
			mod := new(big.Int)
			v.DivMod(v, big.NewInt(23), mod)
			if mod.Cmp(big.NewInt(0)) == 0 {
				return 2
			}
			return 3
		},
		func(v *big.Int) int {
			mod := new(big.Int)
			v.DivMod(v, big.NewInt(19), mod)
			if mod.Cmp(big.NewInt(0)) == 0 {
				return 2
			}
			return 0
		},
		func(v *big.Int) int {
			mod := new(big.Int)
			v.DivMod(v, big.NewInt(13), mod)
			if mod.Cmp(big.NewInt(0)) == 0 {
				return 1
			}
			return 3
		},
		func(v *big.Int) int {
			mod := new(big.Int)
			v.DivMod(v, big.NewInt(17), mod)
			if mod.Cmp(big.NewInt(0)) == 0 {
				return 0
			}
			return 1
		},
	}

	result := part2.Run(orchestrator, 1)
	expected := 2713310158
	if result != expected {
		t.Fatalf("invalid result. got=%d, want=%d\n", result, expected)
	}

}

// the problem here is it is overflowing our integer type (need to use big int?)
