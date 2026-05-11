package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	var output strings.Builder

	words := SplitLine(input)
	for _, word := range words {
		if word == "" {
			for i := 0; i < 8; i++ {
				output.WriteString("\n")
			}
			continue
		}

		generate := RenderLine(word, banner)
		for _, render := range generate {
			output.WriteString(render + "\n")
		}
	}

	return output.String()
}


