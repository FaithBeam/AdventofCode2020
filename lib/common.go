package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(input string) (lines []string) {
	// open file
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func ReadFileInt(input string) (lines []int) {
	// open file
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}
	return
}

func CombineLinesUntilEmptyLine(startIndex int, lines []string) (endIndex int, myString string) {
	myString = ""
	for endIndex = startIndex; endIndex < len(lines); endIndex++ {
		if lines[endIndex] == "" {
			endIndex++
			break
		}
		myString += lines[endIndex] + " "
	}
	return
}

/* Only counts number of empty lines. If you're using this for a for-loop you probably want to add 1 to it so you get the last section of the file. Ex. this returns 4:
abc

a
b
c

ab
ac

a
a
a
a

b
*/
func NumberOfEmptyLines(lines []string) (numBreaks int) {
	for _, line := range lines {
		if line == "" {
			numBreaks += 1
		}
	}
	return
}
