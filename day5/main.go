package main

import (
	"fmt"
	"os"

	// "github.com/josh-weston/advent-of-code-2022/day5/part1"
	"github.com/josh-weston/advent-of-code-2022/day5/part2"
)

func main() {
	f, err := os.OpenFile("day5/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	finalStack := part2.Run(f, 9)
	fmt.Printf("The final arrangement is %q\n", finalStack)

}
