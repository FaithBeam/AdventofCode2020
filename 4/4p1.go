package main

import (
	"strings"

	"../lib"
)

func IsValidPassport(line string) bool {
	if !strings.Contains(line, "ecl") || !strings.Contains(line, "pid") || !strings.Contains(line, "eyr") || !strings.Contains(line, "hcl") || !strings.Contains(line, "byr") || !strings.Contains(line, "iyr") || !strings.Contains(line, "hgt") {
		return false
	}
	return true
}

func main() {
	lines := lib.ReadFile("input.txt")

	// look for number of passports
	numBreaks := 1
	for _, line := range lines {
		if line == "" {
			numBreaks += 1
		}
	}

	validPassports := 0
	curLine := 0
	// foreach passport
	for i := 0; i < numBreaks; i++ {
		currentPassport := ""
		// each passport
		for curLine < len(lines) {
			if lines[curLine] == "" {
				curLine++
				break
			}
			currentPassport += lines[curLine] + " "
			curLine++
		}
		if IsValidPassport(currentPassport) {
			validPassports++
		}
	}
	println(validPassports)
}
