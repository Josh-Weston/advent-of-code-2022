package part2

import (
	"bufio"
	"io"
	"strings"
)

// TODO: start at 'E' and find the first 'a'
func Run(input io.Reader) int {

	grid := [][]byte{}

	// build our 2D grid
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		txt := scanner.Text()
		characters := strings.Split(txt, "")
		row := make([]byte, len(characters))
		for i, c := range characters {
			row[i] = []byte(c)[0]
		}
		grid = append(grid, row)
	}

	// find my starting and end positions
	start := make([]int, 2)
	end := make([]int, 2)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'S' {
				grid[i][j] = 'a'
				start[0] = i
				start[1] = j
			}
			if grid[i][j] == 'E' {
				grid[i][j] = 'z'
				end[0] = i
				end[1] = j
			}
		}
	}

	// Create my graph
	queue := [][]int{start}
	visited := [][]int{start}
	path := [][][]int{[][]int{start}}

	hasVisited := func(node []int) bool {
		for _, v := range visited {
			if v[0] == node[0] && v[1] == node[1] {
				return true
			}
		}
		return false
	}

	canMove := func(val byte, row, col int) bool {
		v := grid[row][col]
		return int(v)-int(val) <= 1 // can climb one, or go down as many as we like; need to change to int() otherwise bytes will overflow
	}

	getAdjacent := func(node []int) [][]int {
		currVal := grid[node[0]][node[1]]
		possibleMoves := [][]int{}
		// can move up
		if node[0] > 0 && canMove(currVal, node[0]-1, node[1]) {
			possibleMoves = append(possibleMoves, []int{node[0] - 1, node[1]})
		}

		// can move down
		if node[0] < len(grid)-1 && canMove(currVal, node[0]+1, node[1]) {
			possibleMoves = append(possibleMoves, []int{node[0] + 1, node[1]})
		}

		// can move left
		if node[1] > 0 && canMove(currVal, node[0], node[1]-1) {
			possibleMoves = append(possibleMoves, []int{node[0], node[1] - 1})
		}

		// can move right
		if node[1] < len(grid[0])-1 && canMove(currVal, node[0], node[1]+1) {
			possibleMoves = append(possibleMoves, []int{node[0], node[1] + 1})
		}
		return possibleMoves
	}

	steps := 0
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		currPath := path[0]
		path = path[1:]

		// see if it is our end
		if v[0] == end[0] && v[1] == end[1] {
			steps = len(currPath) - 1 // ignore first node
			break
		}

		neighbors := getAdjacent(v)
		for _, n := range neighbors {
			if !hasVisited(n) {
				queue = append(queue, n)
				visited = append(visited, n)
				// build my path
				newPath := append(currPath, n)
				path = append(path, newPath)
			}
		}
	}

	return steps
}
