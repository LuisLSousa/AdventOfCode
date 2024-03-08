package main

import (
	"fmt"
	"os"
	"strings"
)

var vowels = map[rune]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
}

var forbiddenStrings = map[string]struct{}{
	"ab": {},
	"cd": {},
	"pq": {},
	"xy": {},
}

const inputFileName = "input.txt"

func main() {
	input, _ := os.ReadFile(inputFileName)

	var niceCounter1, niceCounter2 int
	for _, s := range strings.Fields(string(input)) {
		if isNice1(s) {
			niceCounter1++
		}

		if isNice2(s) {
			niceCounter2++
		}
	}

	fmt.Println("part1: ", niceCounter1)
	fmt.Println("part2: ", niceCounter2)
}

func isNice1(s string) bool {
	var has3Vowels, hasSameLetterTwice, hasForbiddenString bool
	var vowelCounter int
	var previousCharacter rune
	for _, c := range s {
		if c == previousCharacter {
			hasSameLetterTwice = true
		}

		if _, ok := vowels[c]; ok {
			vowelCounter++
		}

		if vowelCounter == 3 {
			has3Vowels = true
		}

		digraph := fmt.Sprintf("%s%s", string(previousCharacter), string(c))
		if _, ok := forbiddenStrings[digraph]; ok {
			hasForbiddenString = true
		}

		previousCharacter = c
	}

	return has3Vowels && hasSameLetterTwice && !hasForbiddenString
}

// s stores the digraph itself, minIndex and maxIndex store the indexes the digraph occupies
type digraph struct {
	minIndex int
	maxIndex int
}

func isNice2(s string) bool {
	var hasSame2LettersTwice, hasSameLetterWith1Between bool
	var previousCharacter, characterBeforeLast rune
	digraphs := make(map[string]digraph)
	for i, c := range s {
		if c == characterBeforeLast {
			hasSameLetterWith1Between = true
		}

		d := fmt.Sprintf("%s%s", string(previousCharacter), string(c))
		v, ok := digraphs[d]
		if !ok {
			digraphs[d] = digraph{
				minIndex: i - 1,
				maxIndex: i,
			}
		}

		if ok && v.maxIndex != i-1 {
			hasSame2LettersTwice = true
		}

		characterBeforeLast = previousCharacter
		previousCharacter = c
	}

	return hasSame2LettersTwice && hasSameLetterWith1Between
}
