package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// The answer is 1-indexed
func main() {
	f, err := os.OpenFile("day1/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	elves := []int{}

	totalCalories := 0
	// This will read 1 line at a time and create the total for each elf
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			panic(err)
		}

		txt := scanner.Text()

		if txt == "" {
			elves = append(elves, totalCalories)
			totalCalories = 0
			continue
		}

		calories, err := strconv.ParseInt(txt, 10, 0)
		if err != nil {
			panic(err)
		}
		totalCalories += int(calories)
	}

	// now that we have the total, we need to determine how many calories the top of elf is carrying

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	total := 0
	for _, v := range elves[:3] {
		total += v
	}

	fmt.Printf("The top 3 elves are have %d calories\n", total)

}
