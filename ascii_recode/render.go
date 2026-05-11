package main

import (
	"fmt"
	"strings"
)

func RenderLine(input string, banner map[rune][]string) []string {
	var builder [8]strings.Builder

	for _, char := range input {
		line, ok := banner[char]
		if !ok {
			line = banner[char]
			continue
		}
		if len(line) != 8 {
			fmt.Printf("warning: character '%c' does not have 8 lines in the banner\n", char)
			continue
		}
		for i := 0; i < 8; i++ {
			builder[i].WriteString(line[i])

		}
	}
	var result []string
	for _, b := range builder {
		result = append(result, b.String())
	}
	return result
}
