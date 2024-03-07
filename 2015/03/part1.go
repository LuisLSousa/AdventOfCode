package main

import (
	"fmt"
	"strings"
)

func part1(input []byte) {
	// grid stores the locations of the houses already visited
	grid := make(map[coordinates]struct{})

	// Santa starts in position 0,0
	var santaPosition coordinates
	grid[santaPosition] = struct{}{}

	for _, s := range strings.Fields(string(input)) {
		for _, r := range s {
			santaPosition = movePosition(santaPosition, r)
			if _, ok := grid[santaPosition]; !ok {
				grid[santaPosition] = struct{}{}
			}

		}
	}

	fmt.Println("Houses visited: ", len(grid))
}
