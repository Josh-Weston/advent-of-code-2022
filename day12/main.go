package main

import (
	"fmt"
	"os"

	"github.com/josh-weston/advent-of-code-2022/day12/part1"
)

func main() {

	f, err := os.Open("day12/input.txt")
	if err != nil {
		panic(err)
	}

	result := part1.BFS(f)
	fmt.Printf("The fewest number of steps is %d\n", result)

}
