package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/LuisLSousa/AdventOfCode/utils"
)

const inputFileName = "input.txt"

// Rule is in the form "a|b" (a comes before b)
type Rule struct {
	a int
	b int
}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rules []Rule
	var update []int
	var part1Sum int
	var totalSum int

	comparison := func(a, b int) int {
		for _, r := range rules {
			if a == r.a && b == r.b {
				// fmt.Println(b, a)
				// fmt.Println(r)
				return -1
			}
		}
		return 0
	}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		switch {
		case len(line) == 2:
			a := utils.MustAtoi(line[0])
			b := utils.MustAtoi(line[1])
			rules = append(rules, Rule{a, b})

		case len(line) == 1 && line[0] != "":
			strUpdate := strings.Split(line[0], ",")
			for _, n := range strUpdate {
				update = append(update, utils.MustAtoi(n))
			}
			if slices.IsSortedFunc(update, comparison) {
				part1Sum += update[len(update)/2]
			}

			slices.SortFunc(update, comparison)
			totalSum += update[len(update)/2]
			update = []int{}
		default:
			continue
		}

	}

	fmt.Println(part1Sum)
	// part2Sum is only the incorrect ones, so total - part1 (the correct ones)
	fmt.Println(totalSum - part1Sum)

	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %s", err)
	}
}
