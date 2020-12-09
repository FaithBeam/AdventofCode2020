package main

import (
	"strconv"
	"strings"

	"../../lib"
)

func main() {
	lines := lib.ReadFile("input.txt")
	// tracks which indexes we've tried changing from jmp -> nop or nop -> jmp
	var changed = make([]int, len(lines))

	acc := 0
	quit := false

	for i := 0; i < len(lines) && !quit; i++ {
		lineCopy := make([]string, len(lines))
		copy(lineCopy, lines)

		for j := 0; j < len(lineCopy); j++ {
			if changed[j] == 0 {
				if strings.Split(lineCopy[j], " ")[0] == "nop" {
					lineCopy[j] = strings.Replace(lineCopy[j], "nop", "jmp", 1)
					changed[j] = 1
					break
				} else if strings.Split(lineCopy[j], " ")[0] == "jmp" {
					lineCopy[j] = strings.Replace(lineCopy[j], "jmp", "nop", 1)
					changed[j] = 1
					break
				}
			}
		}

		// track which instructions we've already visited
		var visited = make([]int, len(lineCopy))
		start := 0
		current := start
		acc = 0

		for current >= 0 && current < len(lineCopy) {
			// Have we visited the current operation
			if visited[current] == 1 {
				break
			}

			visited[current] = 1
			op := strings.Split(lineCopy[current], " ")
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
				current++ // Go to the next line
			}

			if current >= len(lineCopy) {
				quit = true
			}
		}
	}
	println(acc)
}
