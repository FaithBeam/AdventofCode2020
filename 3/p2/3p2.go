package main

import (
	"../../lib"
)

func main() {
	lines := lib.ReadFile("input.txt")
	// move right 1, down 1
	a1 := DoThing(lines, 1, 1)
	a2 := DoThing(lines, 3, 1)
	a3 := DoThing(lines, 5, 1)
	a4 := DoThing(lines, 7, 1)
	a5 := DoThing(lines, 1, 2)
	println(a1 * a2 * a3 * a4 * a5)
}

func DoThing(board []string, hOffset int, vOffset int) int {
	width := len(board[0])
	height := len(board) - 1
	col := 0
	row := 0
	treeCount := 0
	for row < height {
		col = (col + hOffset) % width
		row += vOffset
		if string(board[row][col]) == "#" {
			treeCount += 1
		}
	}
	return treeCount
}
