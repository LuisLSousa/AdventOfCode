package main

import (
	"os"
)

const inputFileName = "input.txt"

type coordinates struct {
	x int
	y int
}

func main() {
	input, _ := os.ReadFile(inputFileName)

	part1(input)
	part2(input)
}

func movePosition(position coordinates, r rune) coordinates {
	switch r {
	case '^':
		position.y++
	case '>':
		position.x++
	case '<':
		position.x--
	case 'v':
		position.y--
	default:
		// any other character
		return position
	}

	return position
}
