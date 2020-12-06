package main

import (
	"os"
	"log"
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Open file
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

	pattern := `^([\d]+)-([\d]+)\s(\w):\s(\w+)$`
	re := regexp.MustCompile(pattern)
	numGoodPasswords := 0 
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		
		min, _ := strconv.Atoi(match[1])
		max, _ := strconv.Atoi(match[2])
		letter := match[3]
		password := match[4]
		cnt := strings.Count(password, letter)
		if min <= cnt && cnt <= max {
			numGoodPasswords++
			println(password)
		}
	}
	println(numGoodPasswords)
}