package main

import "fmt"

func PrintAscii(inputstr string, bannerMap map[rune][]string) {
	for line := 0; line < 8; line++ {
		for _, r := range inputstr {
			fmt.Print(bannerMap[r][line])
		}
	}
}
