package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFileName = "input.txt"

type coordinates struct {
	x, y int
}

/*
	Future Tip: instead of keeping track of the board state, it might be easier to simply track the position and visited places using complex numbers

	type seen map[complex128]struct{}

	s := make(seen)

	pos := 1 + 2i // real number for "x", imaginary number for "y"

	seen[pos] = struct{}{}

	// to change direction, simply multiply by i or -i

	pos *= i // rotate right
	pos *= -i // rotate left


*/

type board map[coordinates]string

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	b := make(board)
	var y int
	var startingPosition coordinates
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for x := range len(line) {
			b[coordinates{x, y}] = line[x]
			if line[x] == "^" {
				startingPosition = coordinates{x, y}
			}
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %s", err)
	}

	fmt.Println("Part 1:", traverseBoard(b, startingPosition))

	var i, j int
	var sum int
	for {
		for {
			v, ok := b[coordinates{i, j}]
			if !ok {
				break
			}

			if v == "X" {
				b[coordinates{i, j}] = "#"
				res := traverseBoard(b, startingPosition)
				if res == -1 {
					sum++
				}
				b[coordinates{i, j}] = "X" // reset it to X
			}
			i++
		}
		i = 0
		j++
		_, ok := b[coordinates{i, j}]
		if !ok {
			break
		}
	}

	fmt.Println("Part 2:", sum)

}

// traverse the board starting from c
func traverseBoard(b board, c coordinates) int {
	b[c] = "X"              // mark the first position as visited
	count, newCount := 1, 0 // start count at 1 to mark the first position
	newCoordinates := c

	// ----- Part 2
	type path struct {
		// start x, y; final x, y
		sx, sy, fx, fy int
	}

	visited := make(map[path]struct{})
	// ------------

	for {
		s := newCoordinates
		// up
		newCount, newCoordinates = traverseUp(b, newCoordinates)
		count += newCount

		f := newCoordinates

		// ----- Part 2
		if _, ok := visited[path{s.x, s.y, f.x, f.y}]; ok {
			return -1
		}
		visited[path{s.x, s.y, f.x, f.y}] = struct{}{}
		// ------------

		if isOut(b, newCoordinates) {
			return count
		}

		// right
		newCount, newCoordinates = traverseRight(b, newCoordinates)
		count += newCount

		if isOut(b, newCoordinates) {
			return count
		}

		// down
		newCount, newCoordinates = traverseDown(b, newCoordinates)
		count += newCount

		if isOut(b, newCoordinates) {
			return count
		}

		// left
		newCount, newCoordinates = traverseLeft(b, newCoordinates)
		count += newCount

		if isOut(b, newCoordinates) {
			return count
		}
	}
}

func traverseUp(b board, c coordinates) (int, coordinates) {
	var count int
	x, y := c.x, c.y-1

loop:
	for {
		newC := coordinates{x, y}
		value, ok := b[newC]
		if !ok {
			return count, coordinates{x, y}
		}

		switch value {
		case "#":
			y++
			break loop
		case ".":
			b[newC] = "X"
			count++
		}
		y--
	}

	return count, coordinates{x, y}
}

func traverseRight(b board, c coordinates) (int, coordinates) {
	var count int
	x, y := c.x+1, c.y

loop:
	for {
		newC := coordinates{x, y}
		value, ok := b[newC]
		if !ok {
			return count, coordinates{x, y}
		}

		switch value {
		case "#":
			x--
			break loop
		case ".":
			b[newC] = "X"
			count++
		}
		x++
	}

	return count, coordinates{x, y}
}

func traverseDown(b board, c coordinates) (int, coordinates) {
	var count int
	x, y := c.x, c.y+1
loop:
	for {
		newC := coordinates{x, y}
		value, ok := b[newC]
		if !ok {
			return count, coordinates{x, y}
		}

		switch value {
		case "#":
			y--
			break loop
		case ".":
			b[newC] = "X"
			count++
		}
		y++
	}

	return count, coordinates{x, y}
}

func traverseLeft(b board, c coordinates) (int, coordinates) {
	var count int
	x, y := c.x-1, c.y

loop:
	for {
		newC := coordinates{x, y}
		value, ok := b[newC]
		if !ok {
			return count, coordinates{x, y}
		}

		switch value {
		case "#":
			x++
			break loop
		case ".":
			b[newC] = "X"
			count++
		}
		x--
	}

	return count, coordinates{x, y}
}

func isOut(b board, c coordinates) bool {
	if _, ok := b[c]; !ok {
		return true
	}
	return false
}
