package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		line := scanner.Text()
		rangesID := strings.Split(line, ",")
		var sumInvalidID int

		for _, str := range rangesID {
			split := strings.Split(str, "-")
			firstID := split[0]
			lastID := split[len(split)-1]

			startID, err := strconv.Atoi(firstID)
			if err != nil {
				return fmt.Errorf("Error converting firstID: %w", err)
			}

			endID, err := strconv.Atoi(lastID)
			if err != nil {
				return fmt.Errorf("Error converting lastID: %w", err)
			}

			for i := startID; i <= endID; i++ {
				toStringID := strconv.Itoa(i)
				if isRepeatedTwice(toStringID) {
					sumInvalidID += i
				}
			}
		}
		fmt.Printf("Sum of all invalid IDs: %d\n", sumInvalidID)
	}
	return nil
}

func isRepeatedTwice(s string) bool {
	length := len(s)

	if length == 0 || length%2 != 0 {
		return false
	}

	mid := length / 2
	firstHalf := s[:mid]
	secondHalf := s[mid:]

	return firstHalf == secondHalf
}
