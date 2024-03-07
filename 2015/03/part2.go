package main

import (
	"fmt"
	"strings"
)

func part2(input []byte) {
	// grid stores the locations of the houses already visited
	grid := make(map[coordinates]struct{})

	// Santa starts in position 0,0
	var santaPosition, robotPosition coordinates
	grid[santaPosition] = struct{}{}

	for _, s := range strings.Fields(string(input)) {
		for i, r := range s {
			switch i % 2 {
			case 0:
				santaPosition = movePosition(santaPosition, r)
			default:
				robotPosition = movePosition(robotPosition, r)
			}

			if _, ok := grid[santaPosition]; !ok {
				grid[santaPosition] = struct{}{}
			}

			if _, ok := grid[robotPosition]; !ok {
				grid[robotPosition] = struct{}{}
			}
		}
	}

	fmt.Println("Houses visited: ", len(grid))
}
