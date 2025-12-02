package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Dial struct {
	position  int
	zeroCount int
}

func main() {
	if err := processInput("input"); err != nil {
		fmt.Println("couldn't process rotations input file")
		os.Exit(1)
	}
}

func processInput(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("couldn't read input file: %w", err)
	}
	defer file.Close()

	dial := Dial{
		position:  50,
		zeroCount: 0,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		direction, quantity, err := parseRotation(scanner.Text())
		if err != nil {
			return fmt.Errorf("couldn't parse rotation: %w", err)
		}

		dial.processRotation(direction, quantity)
	}

	fmt.Printf("Total zero count: %d\n", dial.zeroCount)
	return nil
}

func (dial *Dial) processRotation(direction string, quantity int) {
	for range quantity {
		if direction == "L" {
			dial.position = (dial.position - 1 + 100) % 100
		}
		if direction == "R" {
			dial.position = (dial.position + 1) % 100
		}
		if dial.position == 0 {
			dial.zeroCount++
		}
	}
}

func parseRotation(line string) (string, int, error) {
	direction := string(line[0])
	quantity, err := strconv.Atoi(line[1:])
	if err != nil {
		return "", 0, err
	}

	return direction, quantity, nil
}
