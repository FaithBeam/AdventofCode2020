package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	width := len(lines[0])
	height := len(lines) - 1
	col := 0
	row := 0
	treeCount := 0
	for row < height {
		col = (col + 3) % width
		row += 1
		strArr := []rune(lines[row])
		if string(strArr[col]) == "#" {
			treeCount += 1
		}
	}

	println(treeCount)
}
