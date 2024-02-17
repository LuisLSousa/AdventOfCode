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

var max = map[string]int{
	"r": 12,
	"g": 13,
	"b": 14,
}

func part1() error {
	file, err := os.Open(inputFileName)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var sum int
	for scanner.Scan() {
		sum += parseGame(scanner.Text())
	}

	fmt.Println("Part 1:", sum)
	return nil
}

// parseGame parses a single line from the input file and returns the game ID if the game is valid
func parseGame(s string) int {
	s = strings.Trim(s, "Game ")
	split := strings.Split(s, ":")

	index, rounds := split[0], split[1]
	id, err := strconv.Atoi(index)
	if err != nil {
		log.Printf("input text invalid. Game: %s; error: %s", index, err)
		return 0
	}

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

		if n > max[color] {
			return 0
		}
	}

	return id
}
