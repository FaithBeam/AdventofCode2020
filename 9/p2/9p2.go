// https://adventofcode.com/2020/day/9

package main

import (
	"../../lib"
)

const ninePOneAnswer = 15690279

func main() {
	lines := lib.ReadFileInt("input.txt")
	// this just finds the index of the answer from 9p1 in the lines slice because Go doesn't have a built-in indexOf() method
	answerIndex := lib.SliceIndex(len(lines), func(i int) bool { return lines[i] == ninePOneAnswer })

	// Create slice of successive sums
	sum := 0
	sums := make([]int, answerIndex)
	for i := 0; i < len(sums); i++ {
		sum += lines[i]
		sums[i] = sum
	}

	startIndex := 0
	indexOfSum := -1

	// go through each sum
	for i := 0; i < len(sums); i++ {
		// subtract from each sum
		for j := 0; j < len(sums); j++ {
			sums[j] -= lines[i]
		}

		// If an element in the sums slice contains the answer from 9p1. -1 is false
		indexOfSum = lib.SliceIndex(len(sums), func(i int) bool { return sums[i] == ninePOneAnswer })
		if indexOfSum != -1 {
			startIndex = i + 1
			break
		}
	}

	longestSetOfSums := make([]int, indexOfSum-startIndex+1)
	longestSetOfSums = lines[startIndex : indexOfSum+1]
	println(lib.Min(longestSetOfSums) + lib.Max(longestSetOfSums))
}
