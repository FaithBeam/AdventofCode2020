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

/* taken from https://stackoverflow.com/a/18203895
xs := []int{2, 4, 6, 8}
ys := []string{"C", "B", "K", "A"}
fmt.Println(
    SliceIndex(len(xs), func(i int) bool { return xs[i] == 5 }),
    SliceIndex(len(xs), func(i int) bool { return xs[i] == 6 }),
    SliceIndex(len(ys), func(i int) bool { return ys[i] == "Z" }),
    SliceIndex(len(ys), func(i int) bool { return ys[i] == "A" }))
*/
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// taken from https://stackoverflow.com/a/34259943
func Min(v []int) (m int) {
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return
}

func Max(v []int) (m int) {
	for i, e := range v {
		if i == 0 || e > m {
			m = e
		}
	}
	return
}
