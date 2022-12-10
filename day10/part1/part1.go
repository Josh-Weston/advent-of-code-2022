package part1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// each tick is a cycle
// figure out the signal being sent by the CPU
// register X starts at 1
// register X supports "addx V" = 2 cycles, register is increased by V (V can be negative)
// noop = one cycle, no other effect
// consider the signal strength (cycle number multipled by the value of the X Register) DURING

// (we only increment the register AFTER we move to the next instruction)
// 20, 60, 100, 140, 180, and 220

// start cycle at 1

func Run(input io.Reader) int {

	x := 1
	cycle := 0
	signal := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		txt := scanner.Text()
		cycle++
		if txt == "noop" {
			// we need to register the value
			if cycle <= 220 && ((cycle-20)%40) == 0 {
				fmt.Println(cycle * x)
				signal += (cycle * x)
			}
			continue
		}

		// adding to the register
		fields := strings.Fields(txt)
		v, err := strconv.ParseInt(fields[1], 10, 8)
		if err != nil {
			panic(err)
		}

		// during the cycle
		// we need to register the value
		if cycle <= 220 && ((cycle-20)%40) == 0 {
			fmt.Println(cycle * x)
			signal += (cycle * x)
		}

		cycle++
		if cycle <= 220 && ((cycle-20)%40) == 0 {
			fmt.Println(cycle * x)
			signal += (cycle * x)
		}

		x += int(v)
		continue
	}
	return signal
}
