package part2

import (
	"bufio"
	"io"
	"unicode"
)

// Elves are divided into groups of 3
// The badge is the only item carried by all 3 elves
// at most, two of the Elves will be carrying any other item type
// all of the badges need to be pulled out of the rucksacks, so new badges can be inserted
// every set of 3 lines corresponds to a group

var (
	UPPER_SHIFT int = 38 // -65 + 27
	LOWER_SHIFT int = 96
)

func getLetterPriority(letter rune) int {
	if unicode.IsUpper(letter) {
		return int(letter) - UPPER_SHIFT
	}
	return int(letter) - LOWER_SHIFT
}

func convertToInts(s string) []int {
	converted := make([]int, len(s))
	for i, v := range s {
		converted[i] = getLetterPriority(v)
	}
	return converted
}

func compareSacks(sacks [][]int) int {
	var badge int
	s1, s2, s3 := sacks[0], sacks[1], sacks[2]
	for _, v1 := range s1 {
		found := 0
		for _, v2 := range s2 {
			if v1 == v2 {
				found++
				break
			}
		}
		if found == 1 {
			for _, v3 := range s3 {
				if v1 == v3 {
					found++
					break
				}
			}
		}
		// stop when we find the first value in all three groups
		if found == 2 {
			badge = v1
			break
		}
	}
	return badge
}

func Run(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	total := 0

	group := make([][]int, 3)
	i := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err)
		}
		allItems := scanner.Text()
		group[i] = convertToInts(allItems)
		if i == 2 {
			total += compareSacks(group)
			i = 0
			continue
		}
		i++
	}
	return total
}
