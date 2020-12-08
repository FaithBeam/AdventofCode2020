package main

import (
	"fmt"

	"../../lib"
)

func main() {
	lines := lib.ReadFileInt("input.txt")

	var quit = false
	var e1 = 0
	var e2 = 0
	var e3 = 0
	var result = 0
	for i := 0; i < len(lines) && !quit; i++ {
		for j := 0; j < len(lines) && !quit; j++ {
			for k := 0; k < len(lines) && !quit; k++ {
				if lines[i]+lines[j]+lines[k] == 2020 {
					e1 = lines[i]
					e2 = lines[j]
					e3 = lines[k]
					quit = true
				}
			}
		}
	}
	result = e1 * e2 * e3
	fmt.Printf("%d %d %d\n", e1, e2, e3)
	fmt.Printf("%d", result)
}
