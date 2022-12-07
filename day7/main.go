package main

import (
	"fmt"
	"os"

	// "github.com/josh-weston/advent-of-code-2022/day7/part1"
	"github.com/josh-weston/advent-of-code-2022/day7/part2"
)

func main() {
	f, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	result := part2.Run(f)
	fmt.Printf("The result is %d\n", result)

}
