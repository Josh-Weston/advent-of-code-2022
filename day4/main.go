package main

import (
	"fmt"
	"os"

	"github.com/josh-weston/advent-of-code-2022/day4/part2"
)

func main() {
	f, err := os.OpenFile("day4/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	n := part2.Run(f)
	fmt.Printf("Your total score is %d\n", n)
}
