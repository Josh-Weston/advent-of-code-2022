package main

import (
	"fmt"
	"os"

	"github.com/josh-weston/advent-of-code-2022/day3/part2"
)

func main() {
	f, err := os.OpenFile("day3/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	score := part2.Run(f)
	fmt.Printf("Your total score is %d\n", score)
}
