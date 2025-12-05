package main

import (
	"fmt"
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
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
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

func processLine(line string, d *dial) error {
	if len(line) < 2 {
		return fmt.Errorf("invalid line format: %q", line)
	}
	direction := line[0]
	steps, err := strconv.Atoi(line[1:])
	if err != nil {
		return fmt.Errorf("invalid steps: %w", err)
	}
	d.rotate(direction, steps)
	return nil
}

func run() error {
	lines, err := readFile("input")
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	var d dial
	d.position = 50
	d.size = 100
	numZeroes := 0
	fmt.Println("Starting Position:", d.position)

	for _, line := range lines {
		if line == "" {
			continue
		}
		if err := processLine(line, &d); err != nil {
			fmt.Fprintf(os.Stderr, "warning: skipping line %q: %v\n", line, err)
			continue
		}
		if d.position == 0 {
			numZeroes++
		}
	}

	fmt.Println("Password:", numZeroes)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
