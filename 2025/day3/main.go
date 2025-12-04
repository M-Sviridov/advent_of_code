package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	sumJoltage := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sumJoltage += getLargestJoltage(line)
	}

	fmt.Printf("Total Joltage: %d\n", sumJoltage)

	return nil
}

func getLargestJoltage(line string) int {
	joltage := 0
	highestOne := '0'
	highestTwo := '0'
	var highestOnePos int

	for i, digit := range line {
		if i == len(line)-2 {
			if digit > highestOne {
				highestOne = digit
				highestOnePos = i
				break
			}
			break
		}
		if digit > highestOne {
			highestOne = digit
			highestOnePos = i
		}
	}

	for i := highestOnePos + 1; i < len(line); i++ {
		digit := rune(line[i])
		if digit > highestTwo {
			highestTwo = digit
		}
	}

	joltage = int(highestOne-'0')*10 + int(highestTwo-'0')
	return joltage
}
