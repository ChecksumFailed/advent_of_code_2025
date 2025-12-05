package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type dial struct {
	position int
	size     int
}

func readFile(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(content)
	return strings.Split(fileContent, "\n"), nil
}

func (d *dial) rotate(direction byte, steps int) {
	switch direction {
	case 'R':
		d.position = (d.position + steps) % d.size
	case 'L':
		d.position = (d.position - steps) % d.size
		if d.position < 0 {
			d.position += d.size
		}
	}
}


func processLine(line string, d *dial) {
	direction := line[0]
	steps, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatal(err)
	}
	d.rotate(direction, steps)
}

func run(lines []string) {
	var d dial
	d.position = 50
	d.size = 100
	numZeroes := 0
	fmt.Println("Starting Position:", d.position)

	for _, line := range lines {
		if line == "" {
			continue
		}
		processLine(line, &d)
		if d.position == 0 {
			numZeroes++
		}
	}

	fmt.Println("Password:", numZeroes)
}

func main() {
	lines, err := readFile("input")
	if err != nil {
		log.Fatal(err)
	}
	run(lines)

}
