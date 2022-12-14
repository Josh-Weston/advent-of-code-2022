package main

import (
	"fmt"
	"os"

	// "github.com/josh-weston/advent-of-code-2022/day12/part1"
	"github.com/josh-weston/advent-of-code-2022/day12/part2"
)

func main() {

	f, err := os.Open("day12/input.txt")
	if err != nil {
		panic(err)
	}

	result := part2.Run(f)
	fmt.Printf("The fewest number of steps is %d\n", result)

}
