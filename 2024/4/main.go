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
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for x := range len(line) {
			b[coordinates{x: x, y: y}] = line[x]
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %s", err)
	}

	fmt.Println(searchBoard(b))
}

// returns the count of the word XMAS in the board
func searchBoard(b board) (int, int) {
	var count1, count2 int
	var x, y int
	for {
		_, ok := b[coordinates{x, y}]
		if !ok {
			break // Exit the outer loop when the key does not exist
		}

		for {
			v, ok := b[coordinates{x, y}]
			if !ok {
				x = 0 // reset line coordinate
				break // Exit the inner loop when the key does not exist
			}

			if v == "X" {
				count1 += countXmas(b, coordinates{x, y})
			}

			if v == "A" && isXMas(b, coordinates{x, y}) {
				count2 += 1
			}

			x++

		}
		y++

	}
	return count1, count2
}

// for a given coordinate, search around for the word xmas
func countXmas(b board, c coordinates) int {
	var count int
	// up and back
	if b[coordinates{c.x - 1, c.y - 1}] == "M" && b[coordinates{c.x - 2, c.y - 2}] == "A" && b[coordinates{c.x - 3, c.y - 3}] == "S" {
		count++
	}

	// up
	if b[coordinates{c.x, c.y - 1}] == "M" && b[coordinates{c.x, c.y - 2}] == "A" && b[coordinates{c.x, c.y - 3}] == "S" {
		count++
	}

	// up and front
	if b[coordinates{c.x + 1, c.y - 1}] == "M" && b[coordinates{c.x + 2, c.y - 2}] == "A" && b[coordinates{c.x + 3, c.y - 3}] == "S" {
		count++
	}

	// back
	if b[coordinates{c.x - 1, c.y}] == "M" && b[coordinates{c.x - 2, c.y}] == "A" && b[coordinates{c.x - 3, c.y}] == "S" {
		count++
	}

	// front
	if b[coordinates{c.x + 1, c.y}] == "M" && b[coordinates{c.x + 2, c.y}] == "A" && b[coordinates{c.x + 3, c.y}] == "S" {
		count++
	}

	// down and back
	if b[coordinates{c.x - 1, c.y + 1}] == "M" && b[coordinates{c.x - 2, c.y + 2}] == "A" && b[coordinates{c.x - 3, c.y + 3}] == "S" {
		count++
	}

	// down
	if b[coordinates{c.x, c.y + 1}] == "M" && b[coordinates{c.x, c.y + 2}] == "A" && b[coordinates{c.x, c.y + 3}] == "S" {
		count++
	}

	// down and front
	if b[coordinates{c.x + 1, c.y + 1}] == "M" && b[coordinates{c.x + 2, c.y + 2}] == "A" && b[coordinates{c.x + 3, c.y + 3}] == "S" {
		count++
	}

	return count
}

// for a given coordinate, search around for MAS forming an X (X-MAS)
/*  M   M
      A       (or in any other order with the A in the middle)
	S   S
*/
func isXMas(b board, c coordinates) bool {
	if checkLeftDiagonal(b, c) && checkRightDiagonal(b, c) {
		return true
	}

	return false
}

func checkLeftDiagonal(b board, c coordinates) bool {
	if (b[coordinates{c.x - 1, c.y - 1}] == "M" && b[coordinates{c.x + 1, c.y + 1}] == "S") ||
		(b[coordinates{c.x - 1, c.y - 1}] == "S" && b[coordinates{c.x + 1, c.y + 1}] == "M") {
		return true
	}
	return false
}

func checkRightDiagonal(b board, c coordinates) bool {
	if (b[coordinates{c.x + 1, c.y - 1}] == "M" && b[coordinates{c.x - 1, c.y + 1}] == "S") ||
		(b[coordinates{c.x + 1, c.y - 1}] == "S" && b[coordinates{c.x - 1, c.y + 1}] == "M") {
		return true
	}
	return false
}
