package main

import (
	"regexp"
	"strings"

	"../../lib"
)

func main() {
	lines := lib.ReadFile("input.txt")
	var bags = make(map[string][]string)
	rx := regexp.MustCompile(`((\w+)?\ ?\w+) bags?`)
	for _, line := range lines {
		matches := rx.FindAllStringSubmatch(line, -1)
		for i := 1; i < len(matches); i++ {
			if !strings.Contains(matches[i][1], "no ") {
				bags[matches[0][1]] = append(bags[matches[0][1]], matches[i][1])
			}
		}
	}

	var bagsFlattened = make(map[string]string)
	for k, subBags := range bags {
		for _, bag := range subBags {
			bagsFlattened[k] += one(bags, bag)
		}
	}

	count := 0
	for k := range bagsFlattened {
		if strings.Contains(bagsFlattened[k], "shiny gold") {
			count++
		}
	}
	println(count)
}

func one(bags map[string][]string, current string) (allBags string) {
	if bags[current] == nil {
		return current + ","
	}

	for _, v := range bags[current] {
		allBags += current + "," + one(bags, v)
	}
	return
}
