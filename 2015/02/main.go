package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFileName = "input.txt"

func main() {
	input, _ := os.ReadFile(inputFileName)

	var totalPaper, totalRibbon int
	for _, s := range strings.Fields(string(input)) {
		n := strings.Split(s, "x")
		// assuming that the input is always correct
		n1, _ := strconv.Atoi(n[0])
		n2, _ := strconv.Atoi(n[1])
		n3, _ := strconv.Atoi(n[2])

		totalPaper += calculatePaperNeeded(n1, n2, n3)
		totalRibbon += calculateRibbonNeeded(n1, n2, n3)
	}
	fmt.Println("Paper needed: ", totalPaper)
	fmt.Println("Ribbon needed: ", totalRibbon)
}

func calculatePaperNeeded(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l + min(l*w, l*h, w*h)
}

func calculateRibbonNeeded(l, w, h int) int {
	return l*w*h + 2*l + 2*w + 2*h - 2*max(l, w, h)
}
