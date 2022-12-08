package part1

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
func HiddenInRowOrColumn(startIndex int, row []int) bool {
	currVal := row[startIndex]
	// look to the right
	hiddenRight := false
	for _, v := range row[startIndex+1:] {
		if v >= currVal {
			hiddenRight = true
			break
		}
	}

	hiddenLeft := false
	// look to the left
	for _, v := range row[:startIndex] {
		if v >= currVal {
			hiddenLeft = true
			break
		}
	}
	return hiddenRight && hiddenLeft
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

	totalVisible := 0
	// for each row (excluding top and bottom because we know they are visible)
	for i, v := range rows[1 : len(rows)-1] {
		fmt.Println(v)
		fmt.Println(v[1 : len(v)-1])

		// exclude the starting and ending columns
		values := v[1 : len(v)-1]
		for j := range values {
			hiddenInRow := HiddenInRowOrColumn(j+1, v) // add index to account for skipping first record

			if !hiddenInRow {
				fmt.Printf("value: %d is not hidden in row\n", v[j+1])
				totalVisible++
				continue
			}

			hiddenInColumn := HiddenInRowOrColumn(i+1, cols[j+1]) // index +1 because we skip the first column, i +_1 because we skip the first row
			if !hiddenInColumn {
				fmt.Println("not hidden in column")
				totalVisible++
			}
		}
	}
	return totalVisible + (len(rows) * 2) + ((len(cols) - 2) * 2) // subtract two because the edge columns would be counted by the rows already
}
