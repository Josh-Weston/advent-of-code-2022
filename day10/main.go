package main

import (
	"fmt"
	"os"

	"github.com/josh-weston/advent-of-code-2022/day10/part1"
)

func main() {
	f, err := os.Open("day10/input.txt")
	if err != nil {
		panic(err)
	}
	result := part1.Run(f)
	fmt.Printf("The result is %d: \n", result)
}
