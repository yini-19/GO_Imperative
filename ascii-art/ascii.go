package main

import (
	"fmt"
	"strings"
)

func PrintAscii(inputstr string, bannerMap map[rune][]string) {
	lines := strings.Split(inputstr, "\n")

	for _, r := range lines {
		if r == "" {
			fmt.Println()
		}
		for _, char := range r {
			if _, ok := bannerMap[char]; !ok {
				fmt.Printf("Character '%c' not found in banner map\n", char)
				return
			}
		}
		for line := 0; line <= 8; line++ {
			for _, n := range r {
				fmt.Print(bannerMap[n][line])
			}
			fmt.Println()
		}

	}

}
