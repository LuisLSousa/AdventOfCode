package utils

import (
	"fmt"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func Min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("%q\n", s)
		panic("error during strconv.Atoi")
	}
	return n
}
