package main

import (
	"fmt"
	"strings"
)

func PrintAscii(inputstring string, bannerMap map[rune][]string) {
	lines := strings.Split(inputstring, "\n")
	for _, n := range lines {
		for line := 0; line < 8; line++ {
			for _, r := range n {
				fmt.Print(bannerMap[r][line])
			}
			fmt.Println()
		}
		// if i < len(lines)-1 {
		// 	fmt.Println()
		// }
	}
}
