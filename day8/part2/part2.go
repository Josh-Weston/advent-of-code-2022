package part2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// determine whether there is enough tree cover to keep a tree house hidden
// count the number of trees that are visible
// input is tree height (0=shortest, 9=tallest)
// a tree is visible if all of other trees between it and an edge of the grid are shorter than it.
// only consider trees in the same row or column (up, down, left, and right)
// three trees need to be less than the height of the tree
// we only seem to care about the trees in the middle (not the edges)
// we count the edges as visible, and then the inner pieces

// would like to be able to see a lot of trees
// would like to be able to see a lot of trees
// would like to be able to see a lot of treesPBI

// scenic score, this is where I care about looking out from the middle (need to reverse my ordering when looking to the "left/up")/

// find highest scenic score for ANY tree

func ConvertToIntSlice(s string) []int {
	converted := make([]int, len(s))
	for i, r := range s {
		v, err := strconv.ParseInt(string(r), 10, 64)
		if err != nil {
			panic(err)
		}
		converted[i] = int(v)
	}
	return converted
}

// you need to scan the row from the index of the tree out (you can't just scan the row from left to right) (middle out)
func ScenicScore(startIndex int, row []int) int {
	currVal := row[startIndex]
	// look to the right
	rightDistance := 0
	for i := startIndex + 1; i < len(row); i++ {
		rightDistance++
		v := row[i]
		if v >= currVal {
			break
		}
	}

	leftDistance := 0
	// look to the left
	for i := startIndex - 1; i >= 0; i-- {
		leftDistance++
		v := row[i]
		if v >= currVal {
			break
		}
	}

	return rightDistance * leftDistance
}

// load the input into slices and start walking, (start by looking left and right since it is cheapest)
func Run(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	rows := [][]int{}
	// read them all in
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		txt := scanner.Text()
		rows = append(rows, ConvertToIntSlice(txt))
	}

	// carve out the columns to make it easier
	cols := make([][]int, len(rows[0]))
	for _, row := range rows {
		for j, v := range row {
			cols[j] = append(cols[j], v)
		}
	}

	// for each row (excluding top and bottom because we know they are visible)
	maxScore := 0
	row := 0
	col := 0
	for i, v := range rows {

		for j := range v {
			score := ScenicScore(j, v)       // add index to account for skipping first record
			score *= ScenicScore(i, cols[j]) // index +1 because we skip the first column, i +_1 because we skip the first row
			// fmt.Println("Score: ", score)
			if score > maxScore {
				row = i
				col = j
				maxScore = score
			}
		}
	}
	fmt.Printf("Max score is %d, at row %d x col %d\n", maxScore, row, col)
	return maxScore
}
