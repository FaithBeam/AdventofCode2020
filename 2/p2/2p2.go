package main

import (
	"regexp"
	"strconv"

	"../../lib"
)

func main() {
	lines := lib.ReadFile("input.txt")
	rx := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)$`)
	correctPass := 0

	for _, v := range lines {
		match := rx.FindStringSubmatch(v)
		found := false
		start, _ := strconv.Atoi(match[1])
		start--
		end, _ := strconv.Atoi(match[2])
		end--

		if string(match[4][start]) == match[3] {
			found = !found
		}
		if string(match[4][end]) == match[3] {
			found = !found
		}
		if found {
			correctPass++
		}
	}

	println(correctPass)
}
