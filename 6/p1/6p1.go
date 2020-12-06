package main

import (
	"strings"

	"../distinct"
	"../lib"
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
		currentAnswers = strings.ReplaceAll(currentAnswers, " ", "")
		currentAnswersSplit := strings.Split(currentAnswers, "")
		count += len(distinct.Letters(currentAnswersSplit))
	}
	println(count)
}
