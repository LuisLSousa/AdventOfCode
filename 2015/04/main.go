package main

import (
	"crypto/md5"
	"encoding/hex"
)

const md5key = "bgvyzdsv"

func main() {
	part1()
	part2()
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
