package main

import (
	"bufio"
    "fmt"
    "log"
	"os"
	"strconv"
)

func main() {
    file, err := os.Open("C:\\Users\\TechnoRetardo\\Desktop\\input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	var lines []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		var tmp, _ = strconv.Atoi(scanner.Text())
		lines = append(lines, tmp)
	}
	
	var quit = false
	var e1 = 0
	var e2 = 0
	var result = 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines) && !quit; j++ {
			if lines[i] + lines[j] == 2020 {
				e1 = lines[i]
				e2 = lines[j]
				quit = true
			}
		}
	}
	result = e1 * e2
	fmt.Printf("%d %d\n", e1, e2)
	fmt.Printf("%d", result)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}