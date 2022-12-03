package part1

import (
	"bufio"
	"io"
	"unicode"
)

// each rucksack has two large compartments
// one item type per rucksack
// every item type is identified by a single lowercasse or uppercase letter
// a given rucksack always has the same number of items in each of fits two compartments (split in half)
// a-z = 1-26
// A-Z = 27-52

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

// compareSacks requires the inputs to be sorted
func compareSacks(s1, s2 []int) int {
	total := 0

	// Note: this algorithm will count duplicates
	found := make(map[int]bool)
	for _, v1 := range s1 {
		_, notFound := found[v1]
		if !notFound {
			for _, v2 := range s2 {
				if v1 == v2 {
					total += v1
					found[v1] = true
					break
				}
			}
		}
	}
	return total
}

func convertToInts(s string) []int {
	converted := make([]int, len(s))
	for i, v := range s {
		converted[i] = getLetterPriority(v)
	}
	return converted
}

func Run(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	total := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err)
		}
		allItems := scanner.Text()
		s1 := convertToInts(allItems[:len(allItems)/2])
		s2 := convertToInts(allItems[len(allItems)/2:])

		total += compareSacks(s1, s2)
	}
	return total
}
