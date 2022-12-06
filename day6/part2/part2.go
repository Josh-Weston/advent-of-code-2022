package part2

import (
	"io"
)

// device needs to lock onto their signal
// signal iss a sseries of seemingly-random characters that the device receives one at a time
// subroutine that detects a start-of-packet marker in the datastream.
// start of package = 4 characters that are all different
// index  to end of first 4 character sequence
// 1-index, inclusive of the last character of the 4 character stream

func checkSequence(seq []byte) bool {
	chars := map[byte]bool{}

	for _, b := range seq {
		_, ok := chars[b]
		// duplicate
		if ok {
			return false
		}
		chars[b] = true
	}
	return true
}

func Run(input io.Reader) int {
	allBytes, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}
	startIdx := 0
	endIdx := 14

	found := false
	for l := len(allBytes); endIdx <= l; startIdx, endIdx = startIdx+1, endIdx+1 {
		if checkSequence(allBytes[startIdx:endIdx]) {
			found = true
			break
		}
	}

	if found {
		return endIdx
	}

	return -1
}
