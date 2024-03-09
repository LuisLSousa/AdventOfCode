package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	toggle  command = "toggle"
	turnOn  command = "turn on"
	turnOff command = "turn off"
)

type command string

type coordinates struct {
	x, y int
}

type light struct {
	lit        bool
	brightness int
}

const inputFileName = "input.txt"

func main() {
	file, _ := os.Open(inputFileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// grid holds our lights and a bool stating if they are on (true) or off (false)
	grid := make(map[coordinates]light)

	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), " ")

		switch len(splitString) {
		case 4:
			from := processCoordinates(splitString[1])
			to := processCoordinates(splitString[3])
			grid = processCommand(grid, splitString[0], from, to)
		case 5:
			from := processCoordinates(splitString[2])
			to := processCoordinates(splitString[4])
			grid = processCommand(grid, fmt.Sprintf("%s %s", splitString[0], splitString[1]), from, to)
		default:
			fmt.Println("unexpected instruction length: ", splitString)
		}
	}

	var part1, part2 int
	for _, v := range grid {
		if v.lit {
			part1++
		}
		part2 += v.brightness
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}

func processCoordinates(s string) coordinates {
	coords := strings.Split(s, ",")

	// ignoring errors because we assume the input is always correct
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])

	return coordinates{
		x: x,
		y: y,
	}
}

func processCommand(grid map[coordinates]light, s string, from, to coordinates) map[coordinates]light {
	switch command(s) {
	case toggle:
		grid = toggleLights(grid, from, to)
	case turnOn:
		grid = turnOnLights(grid, from, to)
	case turnOff:
		grid = turnOffLights(grid, from, to)
	default:
		fmt.Println("unknown command: ", s)
	}

	return grid
}

func toggleLights(grid map[coordinates]light, from, to coordinates) map[coordinates]light {
	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			if v, ok := grid[coordinates{x, y}]; ok {
				grid[coordinates{x, y}] = light{!v.lit, v.brightness + 2}
			} else {
				grid[coordinates{x, y}] = light{true, 2}
			}
		}
	}
	return grid

}

func turnOnLights(grid map[coordinates]light, from, to coordinates) map[coordinates]light {
	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			if v, ok := grid[coordinates{x, y}]; ok {
				grid[coordinates{x, y}] = light{true, v.brightness + 1}
			} else {
				grid[coordinates{x, y}] = light{true, 1}
			}
		}
	}
	return grid
}

func turnOffLights(grid map[coordinates]light, from, to coordinates) map[coordinates]light {
	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			if v, ok := grid[coordinates{x, y}]; ok {
				brightness := v.brightness - 1
				if brightness < 0 {
					brightness = 0
				}
				grid[coordinates{x, y}] = light{false, brightness}
			} else {
				grid[coordinates{x, y}] = light{false, 0}
			}
		}
	}
	return grid
}
