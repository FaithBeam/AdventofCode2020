package main

import (
	"strings"

	"../../distinct"
	"../../lib"
)

func main() {
	lines := lib.ReadFile("input.txt")
	groups := lib.NumberOfEmptyLines(lines) + 1

	currentIndex := 0
	currentAnswers := ""
	count := 0
	// Iterate through each group
	for i := 0; i < groups; i++ {
		currentIndex, currentAnswers = lib.CombineLinesUntilEmptyLine(currentIndex, lines)
		currentAnswers = strings.Trim(currentAnswers, " ")
		people := strings.Split(currentAnswers, " ")
		// Get unique answers for each person
		m := distinct.Letters(strings.Split(strings.ReplaceAll(currentAnswers, " ", ""), ""))

		// combine everyone's answers into one string (NOT UNIQUE)
		combinedPeople := strings.Join(people, "")
		// iterate through each unique answer
		for k, _ := range m {
			// if everyone in the group answered the same thing, there should be an equal number of chosen letters to people. So if 4 people chose a, there should be 4 a's
			if strings.Count(combinedPeople, k) == len(people) {
				count++
			}
		}
	}
	println(count)
}
