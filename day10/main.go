package main

import (
	"fmt"
	"os"

	// "github.com/josh-weston/advent-of-code-2022/day10/part1"
	"github.com/josh-weston/advent-of-code-2022/day10/part2"
)

// func main() {
// 	f, err := os.Open("day10/input.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	result := part1.Run(f)
// 	fmt.Printf("The result is %d: \n", result)
// }

func main() {
	f, err := os.Open("day10/input.txt")
	if err != nil {
		panic(err)
	}
	result := part2.Run(f)
	fmt.Println(result[0:40])
	fmt.Println(result[40:80])
	fmt.Println(result[80:120])
	fmt.Println(result[120:160])
	fmt.Println(result[160:200])
	fmt.Println(result[200:240])
}
