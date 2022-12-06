package main

import (
	"fmt"
	"os"

	// "github.com/josh-weston/advent-of-code-2022/day6/part1"
	"github.com/josh-weston/advent-of-code-2022/day6/part2"
)

func main() {
	f, err := os.OpenFile("day6/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := part2.Run(f)
	fmt.Printf("The index of the first sequence is %d\n", result)
}
