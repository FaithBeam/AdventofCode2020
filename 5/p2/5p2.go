package main

import (
	"fmt"
	"sort"

	"../../lib"
)

const Rows = 128
const Cols = 8

func main() {
	lines := lib.ReadFile("input.txt")
	var seatIds []int

	// Loop through each boarding pass
	for _, line := range lines {
		minRows := 0
		maxRows := Rows

		minCols := 0
		maxCols := Cols
		// Split each letter so they're individually indexable
		strArr := []rune(line)

		// Calculate row and column of boarding pass
		for _, letter := range strArr {
			switch string(letter) {
			// Keep bottom half (maxRows changes)
			case "F":
				numRows := (maxRows - minRows) / 2
				maxRows -= numRows
			// Keep upper half (minRows changes)
			case "B":
				numRows := (maxRows - minRows) / 2
				minRows += numRows
			// Keep upper cols (minCols changes)
			case "R":
				numCols := (maxCols - minCols) / 2
				minCols += numCols
			// Keep lower cols (maxCols changes)
			case "L":
				numCols := (maxCols - minCols) / 2
				maxCols -= numCols
			}
		}

		seatIds = append(seatIds, minRows*8+minCols)
	}
	sort.Ints(seatIds)
	cur := seatIds[0]
	next := seatIds[0]
	for i := 0; i < len(seatIds); i++ {
		if next != seatIds[i] {
			fmt.Printf("%d %d", i, next)
			break
		}
		cur = seatIds[i]
		next = cur + 1
	}
}
