package part2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// every section has a unique ID
// each elf is assigned a range of IDs
// overlapping assignments
// 2-4 (inclusive, 2, 3, 4)
// the number that overlap at all

func fullyOverlap(p1Start, p1End, p2Start, p2End int64) bool {
	return (p1Start >= p2Start && p1End <= p2End) || (p2Start >= p1Start && p2End <= p1End)
}

func noOverlap(p1Start, p1End, p2Start, p2End int64) bool {
	return p1Start > p2End || p2Start > p1End
}

func overlap(p1Start, p1End, p2Start, p2End int64) bool {
	return p2Start <= p1End || p2End <= p1Start
}

func Run(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	n := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		pairs := strings.Split(scanner.Text(), ",")
		p1 := strings.Split(pairs[0], "-")
		p2 := strings.Split(pairs[1], "-")

		p1Start, _ := strconv.ParseInt(p1[0], 10, 8)
		p1End, _ := strconv.ParseInt(p1[1], 10, 8)
		p2Start, _ := strconv.ParseInt(p2[0], 10, 8)
		p2End, _ := strconv.ParseInt(p2[1], 10, 8)

		if !noOverlap(p1Start, p1End, p2Start, p2End) && (fullyOverlap(p1Start, p1End, p2Start, p2End) || overlap(p1Start, p1End, p2Start, p2End)) {
			n++
		}
	}
	return n

}
