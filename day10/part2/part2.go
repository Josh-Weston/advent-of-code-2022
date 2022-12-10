package part2

import (
	"bufio"
	"io"
	"math"
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

// X controls the horizontal position of a sprite, sprit is  3 pixels wide, and X sets the horizontal position of the middle of the sprite
// 6 * 40 (6 rows of  40 columns) (0-indexed)
// draw a single pixel during each cycle. Each pixel is a #. We draw DURING the cycle

// determine if the sprite is visible the instante each pixel is drawn. If the sprite is positioned such that one of its three pixels
// is the pixel currently being drawn, the screen produces a lit pixel (#); otherwise, the screen leaves the pixl dark (.)

// what 8 capital letters appear on your CRT?

// the cycles determine which row we are drawing to.
// the CRT draws a single pixel during each cycle

func Run(input io.Reader) []string {

	// this may be easier to do a contiguous array instead of a multidimensional array
	pixels := make([]string, 240)
	for i := 0; i < len(pixels); i++ {
		pixels[i] = "."
	}

	// then when we draw it to the output we break it into the rows
	// 10 (10, 11, 12)
	// 11 (10, 11, 12)
	// 12 (10, 11, 12)
	cycle := -1
	spritePosition := 1 // middle of the sprite

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		txt := scanner.Text()
		cycle++

		row := math.Floor(float64(cycle) / 40)
		rowPosition := spritePosition + int(row*40)

		if txt == "noop" {
			if cycle >= rowPosition-1 && cycle <= rowPosition+1 {
				pixels[cycle] = "#"
			}
			continue
		}

		// adding to the register
		fields := strings.Fields(txt)
		v, err := strconv.ParseInt(fields[1], 10, 8)
		if err != nil {
			panic(err)
		}

		// start of the cycle
		if cycle >= rowPosition-1 && cycle <= rowPosition+1 {
			pixels[cycle] = "#"
		}

		cycle++

		row = math.Floor(float64(cycle) / 40)
		rowPosition = spritePosition + int(row*40)

		// during the cycle
		if cycle >= rowPosition-1 && cycle <= rowPosition+1 {
			pixels[cycle] = "#"
		}

		spritePosition += int(v)
		continue
	}
	return pixels
}
