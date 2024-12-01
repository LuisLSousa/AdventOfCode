package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/LuisLSousa/AdventOfCode/utils"
)

const inputFileName = "input.txt"

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list1, list2 []int

	list2Count := make(map[int]int)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if len(line) != 2 {
			log.Fatalf("Invalid input line: %s", scanner.Text())
		}

		first, err1 := strconv.Atoi(line[0])
		second, err2 := strconv.Atoi(line[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("Error converting to int: %s", scanner.Text())
		}

		list1 = append(list1, first)
		list2 = append(list2, second)
		list2Count[second] += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %s", err)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var part1sum, part2sum int
	for i := 0; i < len(list1); i++ {
		part1sum += utils.Abs(list2[i] - list1[i])

		if c, ok := list2Count[list1[i]]; ok {
			part2sum += list1[i] * c
		}
	}

	fmt.Println("Part 1: ", part1sum)
	fmt.Println("Part 2: ", part2sum)
}
