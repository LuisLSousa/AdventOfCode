package main

import "fmt"

func part1() {
	for i := range 1000000 {
		s := fmt.Sprintf("%s%d", md5key, i)
		md5 := getMD5Hash(s)

		if md5[0:5] == "00000" {
			fmt.Println(i)
			return
		}
	}

	fmt.Println("no number found")
}
