package part1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// the crates need to be rearranged
// crange to rearrange crates so the desired crates are at the top
// the crates move one at a time, but they might move multiple crates in a single instruction
// we are parsing stacks

func parseCrates(s string, n int) []string {
	crateSize := 3 // 3 characters
	stacks := make([]string, n)
	for i := 0; i < n; i++ {
		startIndex := i * 4
		stacks[i] = strings.Trim(strings.Trim(s[startIndex:startIndex+crateSize], " "), "[]")
	}
	return stacks
}

func parseInstruction(s string) []int {
	// 1, 3, and 5 are our instructions
	fields := strings.Fields(s)
	values := make([]int, 3)

	// turn them into integers so we can process them easier
	v, err := strconv.ParseInt(fields[1], 10, 8)
	if err != nil {
		panic(err)
	}
	values[0] = int(v)

	v, err = strconv.ParseInt(fields[3], 10, 8)
	if err != nil {
		panic(err)
	}
	values[1] = int(v)

	v, err = strconv.ParseInt(fields[5], 10, 8)
	if err != nil {
		panic(err)
	}
	values[2] = int(v)

	return values
}

func Run(input io.Reader, n int) string {

	scanner := bufio.NewScanner(input)
	rows := [][]string{}
	numStacks := n
	instructions := false
	stacks := make([][]string, n)

	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		txt := scanner.Text()
		// Instruction separator
		if txt == "" {
			instructions = true

			// 1. I need to turn each column into its own slice
			for i := len(rows) - 1; i >= 0; i-- {
				row := rows[i]
				for j, v := range row {
					if v != "" {
						stacks[j] = append(stacks[j], v)
					}
				}
			}
			fmt.Printf("%+v\n", stacks)
			continue
		}

		if instructions {
			ins := parseInstruction(txt)
			fmt.Println(ins)
			// move
			for i := 0; i < ins[0]; i++ {
				from := ins[1] - 1 // 0 indexed
				to := ins[2] - 1   // 0 indexed

				lastIndex := len(stacks[from]) - 1
				stacks[to] = append(stacks[to], stacks[from][lastIndex])
				stacks[from] = stacks[from][:lastIndex] // take
			}
		} else {
			rows = append(rows, parseCrates(txt, numStacks))
		}
	}

	solution := ""
	for _, col := range stacks {
		l := len(col)
		if l > 0 {
			solution += col[l-1]
		}
	}

	return solution
}
