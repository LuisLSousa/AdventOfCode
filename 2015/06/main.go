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

const inputFileName = "input.txt"

func main() {
	file, _ := os.Open(inputFileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// grid holds our lights and a bool stating if they are on (true) or off (false)
	grid := make(map[coordinates]bool)

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

	var numLitLights int
	for _, v := range grid {
		if v {
			numLitLights++
		}
	}
	fmt.Println("Number of lit lights: ", numLitLights)
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

func processCommand(grid map[coordinates]bool, s string, from, to coordinates) map[coordinates]bool {
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

func toggleLights(grid map[coordinates]bool, from, to coordinates) map[coordinates]bool {
	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			if v, ok := grid[coordinates{x, y}]; ok {
				grid[coordinates{x, y}] = !v
			} else {
				grid[coordinates{x, y}] = true
			}
		}
	}
	return grid

}

func turnOnLights(grid map[coordinates]bool, from, to coordinates) map[coordinates]bool {
	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			grid[coordinates{x, y}] = true
		}
	}
	return grid
}

func turnOffLights(grid map[coordinates]bool, from, to coordinates) map[coordinates]bool {
	for x := from.x; x <= to.x; x++ {
		for y := from.y; y <= to.y; y++ {
			grid[coordinates{x, y}] = false
		}
	}
	return grid
}
