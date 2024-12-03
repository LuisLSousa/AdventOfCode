package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const inputFileName = "input.txt"

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	m1 := regexp.MustCompile(`mul\(\d+,\d+\)`)
	m2 := regexp.MustCompile(`[0-9]+`)

	// part 2 m1
	p2m1 := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)

	var part1Sum, part2Sum int
	flag := true

	for scanner.Scan() {
		muls := m1.FindAllString(scanner.Text(), -1)
		p2 := p2m1.FindAllString(scanner.Text(), -1)

		for _, s := range muls {
			n := m2.FindAllString(s, 2)
			first, _ := strconv.Atoi(n[0])
			second, _ := strconv.Atoi(n[1])
			part1Sum += first * second
		}

		for _, s := range p2 {
			switch {
			case s == "do()":
				flag = true
			case s == "don't()":
				flag = false
			default:
				if flag {
					n := m2.FindAllString(s, 2)
					first, _ := strconv.Atoi(n[0])
					second, _ := strconv.Atoi(n[1])
					part2Sum += first * second
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %s", err)
	}

	fmt.Println(part1Sum)
	fmt.Println(part2Sum)
}
