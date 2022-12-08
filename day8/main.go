package main

import (
	"fmt"
	"os"

	"github.com/josh-weston/advent-of-code-2022/day8/part2"
)

func main() {
	f, err := os.Open("day8/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := part2.Run(f)
	fmt.Printf("Highest scenic score possible is: %d\n", result)
}
