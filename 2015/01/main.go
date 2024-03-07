package main

import (
	"fmt"
	"os"
	"strings"
)

const inputFileName = "input.txt"

func main() {
	input, _ := os.ReadFile(inputFileName)

	var floor int
	basementIndex := -1
	var basementFound bool
	for _, s := range strings.Fields(string(input)) {
		for i, r := range s {
			if r == '(' {
				floor += 1
			}
			if r == ')' {
				floor -= 1
			}
			if floor == -1 && !basementFound {
				basementIndex = i + 1
				basementFound = true
			}
		}
	}
	fmt.Println("Final floor: ", floor)
	fmt.Println("Basement index: ", basementIndex)
}
