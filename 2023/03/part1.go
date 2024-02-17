package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type coordinates struct {
	x, y int
}

func part1() error {
	input, _ := os.ReadFile(inputFileName)

	// the grid will store numbers in the coordinates they appear in
	grid := make(map[coordinates]int)

	// store coordinates of characters
	var characterPositions []coordinates

	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r != '.' && !unicode.IsDigit(r) {
				characterPositions = append(characterPositions, coordinates{x, y})
			}

			if unicode.IsDigit(r) {
				grid[coordinates{x, y}] = int(r - '0')
			}
		}
	}

	partNumberCoordinates := make(map[coordinates]struct{})
	for _, c := range characterPositions {

		// merge the maps
		for k, v := range findPartNumberCoordinates(grid, c.x, c.y) {
			partNumberCoordinates[k] = v
		}
	}

	// iterate over entire list again, for each number found that has coordinates in partNumbersCoordinates, sum it
	var sum, currentNumber int
	var isAdjacent bool
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if _, ok := partNumberCoordinates[coordinates{x: x, y: y}]; ok {
				isAdjacent = true
			}

			if unicode.IsDigit(r) {
				currentNumber = combineNumbers(currentNumber, int(r-'0'))
				fmt.Println(currentNumber)

			} else if !isAdjacent { // not a digit and is not adjacent
				currentNumber = 0

			} else { // not a digit and isAdjacent
				sum += currentNumber
				currentNumber = 0
				isAdjacent = false
			}

		}
	}

	fmt.Println("Part 1: ", sum)
	return nil
}

func findPartNumberCoordinates(grid map[coordinates]int, x, y int) map[coordinates]struct{} {
	partCoordinates := make(map[coordinates]struct{})
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if _, ok := grid[coordinates{i, j}]; ok {
				partCoordinates[coordinates{i, j}] = struct{}{}
			}
		}
	}

	return partCoordinates
}

func combineNumbers(a, b int) int {
	return a*10 + b
}
