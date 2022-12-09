package part2

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

// model the ropes on a 2D grid.
// head (H) and tail (T) must always be touching (diagonally adjacent and overlapping count as touching)

// if not touching, and not in the same row or column, the tail moves diagonally to keep up
// H and T start at the same position
// after each step, update the position of the tail
// the tail just needs to be in the last position of the H? Not really, T doesn't move if it is able to "touch" H
// how many positions does the tail of the rope visit at least once?

// isTouching()
// inSameRow()
// inSameColumn()
// move()

// the head is mapping out the environment as we go?
// we can use a relative reference from the start to determine where tail has travelled

// this is like a game of snake, but the tail only needs to move if it is not within reach of the last knot (the 8th)
// the much of the same logic applies, we just need to determine where the second to last knot is?

type Coords struct {
	x int
	y int
}

type Head struct {
	Coords
}

// TODO: this is the logic required for the Head
func (h *Head) Move(ins string, dir int) {
	switch ins {
	case "R":
		h.x += dir
	case "L":
		h.x -= dir
	case "U":
		h.y += dir
	case "D":
		h.y -= dir
	default:
		panic("Unknown direction provided")
	}
}

type Tail struct {
	Coords
	Visited map[string]bool
}

// TODO: this is the logic required for the Tail, it receives the x, y coordinates for where the head is
func (t *Tail) Follow(x, y int) {

	// it needs to move
	if t.Distance(t.x, t.y, x, y) >= 2 {
		// in the same column, need to change the row
		/** MOVE UD **/
		if t.x == x {
			if y > t.y {
				t.y++
			} else {
				t.y--
			}
			t.Visited[fmt.Sprintf("%d,%d", t.x, t.y)] = true
			return
		}

		// in the same row, need to change the column
		/** MOVE LR **/
		if t.y == y {
			if x > t.x {
				t.x++
			} else {
				t.x--
			}
			t.Visited[fmt.Sprintf("%d,%d", t.x, t.y)] = true
			return
		}

		/** MOVE DIAGONALLY **/
		if y > t.y {
			t.y++
		} else {
			t.y--
		}

		// in the same row, need to change the column
		if x > t.x {
			t.x++
		} else {
			t.x--
		}
		t.Visited[fmt.Sprintf("%d,%d", t.x, t.y)] = true
		return
	}
}

func (t *Tail) Distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

func Run(input io.Reader) int {

	scanner := bufio.NewScanner(input)

	// both start at the origin
	start := Coords{x: 0, y: 0}
	H := Head{Coords: start}

	numFollowers := 9
	followers := make([]Tail, numFollowers)
	for i := 0; i < numFollowers; i++ {
		followers[i] = Tail{Coords: start, Visited: map[string]bool{"0,0": true}}
	}

	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		txt := scanner.Text()
		fields := strings.Fields(txt)
		ins := fields[0]
		units, err := strconv.ParseInt(fields[1], 10, 8)
		if err != nil {
			panic(err)
		}

		for i := 0; i < int(units); i++ {
			H.Move(ins, 1)

			// This is like a pipeline
			followers[0].Follow(H.x, H.y)
			for i := 1; i < len(followers); i++ {
				followers[i].Follow(followers[i-1].x, followers[i-1].y)
			}
		}

		if err != nil {
			panic(err)
		}
	}
	return len(followers[len(followers)-1].Visited)
}
