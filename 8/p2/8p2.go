package main

import (
	"strconv"
	"strings"

	"../../lib"
)

func main() {
	lines := lib.ReadFile("input.txt")
	var visited = make([]int, len(lines))
	start := 0
	current := start
	acc := 0

	for current >= 0 && current < len(lines) {
		// Have we visited the current operation
		if visited[current] == 1 {
			break
		}

		visited[current] = 1
		op := strings.Split(lines[current], " ")
		if string(op[1][0]) == "+" {
			op[1] = strings.TrimPrefix(op[1], "+")
		}
		val, _ := strconv.Atoi(op[1])
		switch op[0] {
		case "acc":
			acc += val
			current++ // Go to the next line
		case "jmp":
			current += val
		case "nop":
			acc += val
			current++ // Go to the next line
		}
	}
	println(acc)
}
