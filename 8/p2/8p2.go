package main

import (
	"strconv"
	"strings"

	"../../lib"
)

// unfinished and probably going to rewrite
func main() {
	lines := lib.ReadFile("input.txt")
	// tracks which instructions we visit
	var visited = make([]int, len(lines))
	// tracks which instructions we visit when we fix an instruction
	var subVisited = make([]int, len(lines))
	// tracks which instruction we fixed so we can go back if it's incorrect
	branched := 0
	changedOp := false
	current := 0
	//accumulator
	acc := 0
	// accumulator for fixed instruction
	subAcc := 0

	for current >= 0 && current < len(lines) {
		if changedOp == true {
			// jump out of sub route
			if visited[current] == 1 || subVisited[current] == 1 {
				current = branched + 1
				subAcc = 0
				subVisited = make([]int, len(lines))
				changedOp = false
			}
		}

		// Mark instruction as visited depending if it's a sub route
		if changedOp == true {
			subVisited[current] = 1
		} else {
			visited[current] = 1
		}

		// split the line into instruction and value
		op := strings.Split(lines[current], " ")
		// remove the + sign from the value
		if string(op[1][0]) == "+" {
			op[1] = strings.TrimPrefix(op[1], "+")
		}

		// get value of the instruction
		val, _ := strconv.Atoi(op[1])

		// if we haven't fixed an instruction yet, fix it
		if changedOp == false {
			branched = current
			if op[0] == "nop" {
				op[0] = "jmp"
				changedOp = true
			} else if op[0] == "jmp" {
				op[0] = "nop"
				changedOp = true
			}

			switch op[0] {
			case "acc":
				acc += val
				current++ // Go to the next line
			case "jmp":
				if val == 0 {
					current++
				} else {
					current += val
				}
			case "nop":
				acc += val
				current++ // Go to the next line
			}
		} else {
			switch op[0] {
			case "acc":
				subAcc += val
				current++ // Go to the next line
			case "jmp":
				if val == 0 {
					current++
				} else {
					current += val
				}
			case "nop":
				subAcc += val
				current++ // Go to the next line
			}
		}
	}
	println(acc + subAcc)
}
