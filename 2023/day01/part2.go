package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// will be used to compare the strings found with existing numbers
var (
	threeLetterNumbers = []string{"one", "two", "six"}
	fourLetterNumbers  = []string{"four", "five", "nine"}
	fiveLetterNumbers  = []string{"three", "seven", "eight"}

	stringToDigit = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func part2() error {
	file, err := os.Open(inputFileName)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var sum int
	for scanner.Scan() {
		sum += findCalibrationValuesInString2(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("Part 2: ", sum)
	return nil
}

func findCalibrationValuesInString2(s string) int {
	first := findFirstNumberInString(s, false)

	last := findFirstNumberInString(reverseString(s), true)

	calibrationNumber := first*10 + last
	return calibrationNumber
}

// reversed is used to parse written numbers in strings that have been reversed
func findFirstNumberInString(s string, reversed bool) int {
	var threeLetter, fourLetter, fiveLetter string

	runeString := []rune(s)
	stringSize := len(runeString)

	// find first number in string
	for i := 0; i < stringSize; i++ {
		if unicode.IsDigit(runeString[i]) {
			return int(runeString[i] - '0')

		}

		// possible string numbers
		if i+3 <= stringSize {
			threeLetter = string(runeString[i : i+3])
			if reversed {
				threeLetter = reverseString(threeLetter)
			}
		}

		if i+4 <= stringSize {
			fourLetter = string(runeString[i : i+4])
			if reversed {
				fourLetter = reverseString(fourLetter)
			}
		}

		if i+5 <= stringSize {
			fiveLetter = string(runeString[i : i+5])
			if reversed {
				fiveLetter = reverseString(fiveLetter)
			}
		}

		switch {
		case i+3 <= stringSize && isThreeLetterNumber(threeLetter):
			return stringToDigit[threeLetter]
		case i+4 <= stringSize && isFourLetterNumber(fourLetter):
			return stringToDigit[fourLetter]
		case i+5 <= stringSize && isFiveLetterNumber(fiveLetter):
			return stringToDigit[fiveLetter]
		default:
			continue
		}
	}

	return -1
}

func reverseString(s string) string {
	runeString := []rune(s)
	length := len(runeString)
	for i := 0; i < length/2; i++ {
		runeString[i], runeString[length-1-i] = runeString[length-1-i], runeString[i]
	}

	return string(runeString)
}

func isThreeLetterNumber(s string) bool {
	for _, v := range threeLetterNumbers {
		if s == v {
			return true
		}
	}

	return false
}

func isFourLetterNumber(s string) bool {
	for _, v := range fourLetterNumbers {
		if s == v {
			return true
		}
	}

	return false
}

func isFiveLetterNumber(s string) bool {
	for _, v := range fiveLetterNumbers {
		if s == v {
			return true
		}
	}

	return false
}
