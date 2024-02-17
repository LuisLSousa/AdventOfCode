package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
		sum += parseGame2(scanner.Text())
	}

	fmt.Println("Part 2:", sum)
	return nil
}

// parseGame parses a single line from the input file and returns the game ID if the game is valid
func parseGame2(s string) int {
	// minimum cubes needed
	cubes := make(map[string]int)
	cubes["r"] = 0
	cubes["g"] = 0
	cubes["b"] = 0

	rounds := strings.Split(s, ":")[1]

	r := regexp.MustCompile(`\d+\s[a-z]{1}`)
	games := r.FindAll([]byte(rounds), -1)
	for _, game := range games {
		splitGame := strings.Split(string(game), " ")

		color := splitGame[1]

		n, err := strconv.Atoi(splitGame[0])
		if err != nil {
			log.Printf("input text invalid. %s %s cubes found: %s", splitGame[0], color, err)
			return 0
		}

		if n > cubes[color] {
			cubes[color] = n
		}

	}

	return cubes["r"] * cubes["g"] * cubes["b"]
}
