package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/LuisLSousa/AdventOfCode/utils"
)

const inputFileName = "input.txt"

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Fatalf("Invalid line format: %s", line)
		}

		targetStr := strings.TrimSpace(parts[0])
		target, err := strconv.Atoi(targetStr)
		if err != nil {
			log.Fatalf("Invalid target number '%s' in line: %s", targetStr, line)
		}

		numsStr := strings.TrimSpace(parts[1])
		numsStrSlice := strings.Fields(numsStr)
		if len(numsStrSlice) == 0 {
			continue
		}

		numbers := make([]int, len(numsStrSlice))
		for i, s := range numsStrSlice {
			numbers[i] = utils.MustAtoi(s)
		}

		currentResults := make(map[int]struct{})
		currentResults[numbers[0]] = struct{}{}

		for _, num := range numbers[1:] {
			nextResults := make(map[int]struct{})
			// for each new number found, go over all the results found previously and add both the sum and product to the set of results
			for val := range currentResults {
				sum := val + num
				nextResults[sum] = struct{}{}

				product := val * num
				nextResults[product] = struct{}{}

				// [PART 2]: Uncomment the following two lines
				// concatenation := utils.MustAtoi(strconv.Itoa(val) + strconv.Itoa(num))
				// nextResults[concatenation] = struct{}{}
			}
			currentResults = nextResults
		}

		// if the target is found on the set of results, sum it to the sum
		if _, ok := currentResults[target]; ok {
			sum += target
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %s", err)
	}

	fmt.Println("Sum:", sum)
}
