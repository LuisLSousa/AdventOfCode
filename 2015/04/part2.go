package main

import "fmt"

func part2() {
	for i := range 10000000 {
		s := fmt.Sprintf("%s%d", md5key, i)
		md5 := getMD5Hash(s)

		if md5[0:6] == "000000" {
			fmt.Println(i)
			return
		}
	}

	fmt.Println("no number found")
}
