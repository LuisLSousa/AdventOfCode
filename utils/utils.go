package utils

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
