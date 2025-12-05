package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(content)
	return strings.Split(fileContent, "\n"), nil
}

func rotateDial(line string, curPosition int) int {
	direction := line[0]
	steps, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatal(err)
	}

	switch direction {
	case 'R':
		curPosition = (curPosition + steps) % 100
	case 'L':
		curPosition = (curPosition - steps) % 100
		if curPosition < 0 {
			curPosition += 100
		}
	}
	return curPosition
}

func main() {
	curPosition := 50
	numZeroes := 0
	fmt.Println("Starting Position:", curPosition)

	lines, err := readFile("input")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		curPosition = rotateDial(line, curPosition)
		fmt.Println("Rotated ", line, " New Position:", curPosition)
		if curPosition == 0 {
			numZeroes++
		}
	}

	fmt.Println("Number of times position was zero:", numZeroes)
}
