// https://adventofcode.com/2020/day/9

package main

import (
	"../../lib"
)

const preamble = 25

func main() {
	lines := lib.ReadFileInt("input.txt")
	history := make([]int, preamble)
	lastIndex := 0

	for lastIndex < len(lines) {
		// shift the current x numbers
		history = lines[lastIndex : lastIndex+preamble]
		if !Process(history, lines[lastIndex+preamble]) {
			break
		}
		lastIndex++
	}
	println(lines[lastIndex+preamble])
}

// Does this slice contain two numbers that sum to some number
// this most certainly has an out of bounds exception in the inner for-loop
func Process(history []int, find int) bool {
	for i := 0; i < len(history); i++ {
		for j := i + 1; j < len(history); j++ {
			if history[i]+history[j] == find {
				return true
			}
		}
	}
	return false
}
