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
	if err := processRotationsInput("input"); err != nil {
		fmt.Println("couldn't process rotations input file")
		os.Exit(1)
	}
}

func processRotationsInput(filename string) error {
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

		dial.rotate(direction, quantity)
		dial.checkAndCountZero()
	}
	fmt.Printf("Total zero count: %d\n", dial.zeroCount)
	return nil
}

func parseRotation(line string) (string, int, error) {
	direction := string(line[0])
	quantity, err := strconv.Atoi(line[1:])
	if err != nil {
		return "", 0, err
	}

	return direction, quantity, nil
}

func (dial *Dial) rotate(direction string, quantity int) {
	if direction == "L" {
		dial.rotateLeft(quantity)
	}

	if direction == "R" {
		dial.rotateRight(quantity)
	}
}

func (dial *Dial) rotateLeft(quantity int) {
	dial.position = ((dial.position-quantity)%100 + 100) % 100
}

func (dial *Dial) rotateRight(quantity int) {
	dial.position = (dial.position + quantity) % 100
}

func (dial *Dial) checkAndCountZero() {
	if dial.position == 0 {
		dial.zeroCount++
	}
}
