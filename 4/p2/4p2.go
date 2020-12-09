package main

import (
	"strconv"
	"strings"
	"unicode"

	"../../lib"
)

func main() {
	lines := lib.ReadFile("input.txt")
	numGroups := lib.NumberOfEmptyLines(lines) + 1

	groupIndex := 0
	line := ""
	count := 0
	fields := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for i := 0; i < numGroups; i++ {
		groupIndex, line = lib.CombineLinesUntilEmptyLine(groupIndex, lines)
		if ValidOne(line, fields) {
			if ValidTwo(line) {
				count++
			}
		}
	}
	println(count)
}

func ValidOne(line string, fields [7]string) bool {
	for _, v := range fields {
		if !strings.Contains(line, v) {
			return false
		}
	}
	return true
}

func ValidTwo(line string) bool {
	split := strings.Split(line, " ")
	if split[len(split)-1] == "" {
		split = split[:len(split)-1]
	}
	passPortMap := make(map[string]string)
	for _, v := range split {
		passPortMap[v[:3]] = v[4:]
	}

	if atoi(passPortMap["byr"]) < 1920 || atoi(passPortMap["byr"]) > 2002 || len(passPortMap["byr"]) != 4 {
		return false
	}

	if atoi(passPortMap["iyr"]) < 2010 || atoi(passPortMap["iyr"]) > 2020 || len(passPortMap["iyr"]) != 4 {
		return false
	}

	if atoi(passPortMap["eyr"]) < 2020 || atoi(passPortMap["eyr"]) > 2030 || len(passPortMap["eyr"]) != 4 {
		return false
	}

	if strings.HasSuffix(passPortMap["hgt"], "cm") {
		hgt := atoi(passPortMap["hgt"][:3])
		if hgt < 150 || hgt > 193 {
			return false
		}
	} else if strings.HasSuffix(passPortMap["hgt"], "in") {
		hgt := atoi(passPortMap["hgt"][:2])
		if hgt < 59 || hgt > 76 {
			return false
		}
	} else {
		return false
	}

	if string(passPortMap["hcl"][0]) == "#" || len(passPortMap["hcl"]) == 7 {
		for i, v := range passPortMap["hcl"] {
			if i == 0 {
				continue
			} else if !funnyString(string(v)) && !unicode.IsNumber(v) {
				return false
			}
		}
	} else {
		return false
	}

	if passPortMap["ecl"] == "amb" || passPortMap["ecl"] == "blu" || passPortMap["ecl"] == "brn" || passPortMap["ecl"] == "gry" || passPortMap["ecl"] == "grn" || passPortMap["ecl"] == "hzl" || passPortMap["ecl"] == "oth" {

	} else {
		return false
	}

	if len(passPortMap["pid"]) == 9 {
		for _, v := range passPortMap["pid"] {
			if !unicode.IsNumber(v) {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func funnyString(strang string) bool {
	if strang >= "a" && strang <= "f" {
		return true
	}
	return false
}
