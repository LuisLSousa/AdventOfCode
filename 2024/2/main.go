package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFileName = "input.txt"

func isSafe(row []string) bool {
	inc := make([]int, len(row)-1)
	for i := 0; i < len(row)-1; i++ {
		curr, _ := strconv.Atoi(row[i])
		next, _ := strconv.Atoi(row[i+1])
		inc[i] = next - curr
	}

	increasingSet := true
	decreasingSet := true
	for _, v := range inc {
		if v < 1 || v > 3 {
			increasingSet = false
		}
		if v > -1 || v < -3 {
			decreasingSet = false
		}
	}

	return increasingSet || decreasingSet
}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var part1Sum, part2Sum int
	for scanner.Scan() {
		report := strings.Fields(scanner.Text())
		if isSafe(report) {
			part1Sum++
		}

		// for each report, create all possible subreports without one of the numbers and run part 1 on the subreport
		// if any of them isSafe, immediately increment part2Sum and break (so we don't get multiple safe counts for a single report)
		for i := range report {
			newReport := append([]string{}, report[:i]...)
			newReport = append(newReport, report[i+1:]...)
			if isSafe(newReport) {
				part2Sum++
				break
			}
		}
	}

	fmt.Println(part1Sum)
	fmt.Println(part2Sum)
}
