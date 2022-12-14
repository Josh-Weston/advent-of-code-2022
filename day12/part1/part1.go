package part1

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"strings"
)

func Run(input io.Reader, nPermutations int) int {

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

	type Message struct {
		Steps  int
		Failed bool
	}

	best := math.MaxInt64
	msgChan := make(chan Message)
	doneChan := make(chan struct{})
	defer close(msgChan)
	defer close(doneChan)

	canMove := func(val byte, row, col int) bool {
		v := grid[row][col]
		return v-val <= 1
	}

	samePosition := func(p1, p2 []int) bool {
		return p1[0] == p2[0] && p1[1] == p2[1]
	}

	// our search function
	search := func(done <-chan struct{}, m chan<- Message) {
		currPos := []int{start[0], start[1]}
		prevPos := []int{start[0], start[1]}
		currVal := byte('a') // the start shows as "s", but represents an "a"

		steps := 0
		maxSteps := 1_000
		for {

			select {
			case <-done:
				return
			default:

				if samePosition(currPos, end) {
					m <- Message{
						Steps:  steps,
						Failed: false,
					}
					return
				}

				if steps == maxSteps {
					m <- Message{
						Steps:  steps,
						Failed: true,
					}
					return
				}

				possibleMoves := [][]int{}

				// TODO: this is likely not checking for "E" properly
				// can move up
				if currPos[0] > 0 && canMove(currVal, currPos[0]-1, currPos[1]) && currPos[0]-1 != prevPos[0] {
					possibleMoves = append(possibleMoves, []int{currPos[0] - 1, currPos[1]})
				}

				// can move down
				if currPos[0] < len(grid)-1 && canMove(currVal, currPos[0]+1, currPos[1]) && currPos[0]+1 != prevPos[0] {
					possibleMoves = append(possibleMoves, []int{currPos[0] + 1, currPos[1]})
				}

				// can move left
				if currPos[1] > 0 && canMove(currVal, currPos[0], currPos[1]-1) && currPos[1]-1 != prevPos[1] {
					possibleMoves = append(possibleMoves, []int{currPos[0], currPos[1] - 1})
				}

				// can move right
				if currPos[1] < len(grid[0])-1 && canMove(currVal, currPos[0], currPos[1]+1) && currPos[1]+1 != prevPos[1] {
					possibleMoves = append(possibleMoves, []int{currPos[0], currPos[1] + 1})
				}

				// if we can't make any moves, then kill it
				if len(possibleMoves) == 0 {
					m <- Message{
						Steps:  steps,
						Failed: true,
					}
					return
				}
				// randomly pick a move
				move := possibleMoves[rand.Intn(len(possibleMoves))]
				fmt.Printf("%+v\n", move)
				copy(prevPos, currPos)
				copy(currPos, move)
				currVal = grid[currPos[0]][currPos[1]]

				steps++
			}

		}

	}

	for i := 0; i < 1; i++ {
		go search(doneChan, msgChan)
	}

	n := 0

	for v := range msgChan {
		if !v.Failed {
			if v.Steps < best {
				best = v.Steps
				n = 1
				break
			} else {
				n++
			}
			if n == nPermutations {
				doneChan <- struct{}{}
				break
			}
		}
		go search(doneChan, msgChan)
	}

	return best
}

/* It looks like the real solution will use BFS or Dijkstra's algorithm */

/* BFS will allow us to know where the dead-ends are instead of creating random walks */
func BFS(input io.Reader) int {

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
