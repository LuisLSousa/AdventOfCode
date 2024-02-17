package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func part1() error {
	file, err := os.Open(inputFileName)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var sum int
	for scanner.Scan() {
		sum += findCalibrationValuesInString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("Part 1: ", sum)
	return nil
}

func findCalibrationValuesInString(s string) int {
	runeString := []rune(s)
	var first, last int
	// find first number in string
	for _, r := range runeString {
		if unicode.IsDigit(r) {
			first = int(r - '0')
			break
		}
	}

	// find last number in string
	for i := len(runeString) - 1; i >= 0; i-- {
		if unicode.IsDigit(runeString[i]) {
			last = int(runeString[i] - '0')
			break
		}
	}

	calibrationNumber := first*10 + last

	return calibrationNumber
}
