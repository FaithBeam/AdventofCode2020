package main

import (
	"regexp"
	"strconv"
	"strings"

	"../../lib"
)

// This program runs for like 20 minutes fyi
func main() {
	lines := lib.ReadFile("input.txt")
	var bags = make(map[string][]string)
	rx := regexp.MustCompile(`(\d+)?\s?((\w+)?\ ?\w+) bags?`)
	for _, line := range lines {
		matches := rx.FindAllStringSubmatch(line, -1)
		for i := 1; i < len(matches); i++ {
			if !strings.Contains(matches[i][2], "no ") {
				if matches[i][1] != "" {
					amt, _ := strconv.Atoi(matches[i][1])
					for repeat := 0; repeat < amt; repeat++ {
						bags[matches[0][2]] = append(bags[matches[0][2]], matches[i][2])
					}
				} else {
					bags[matches[0][2]] = append(bags[matches[0][2]], matches[i][2])
				}
			}
		}
	}

	var bagsFlattened = make(map[string]int)
	for k, subBags := range bags {
		for _, bag := range subBags {
			bagsFlattened[k] += one(bags, bag)
		}
	}
	println(bagsFlattened["shiny gold"])
}

func one(bags map[string][]string, current string) (allBags int) {
	allBags++

	if bags[current] == nil {
		return
	}

	for _, v := range bags[current] {
		allBags += one(bags, v)
	}

	return
}
