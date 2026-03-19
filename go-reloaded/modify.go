package main

import (
	"strconv"
	"strings"
)

func ProcessText(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		word := words[i]

		switch {

		case word == "(hex)":
			value, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = strconv.Itoa(int(value))
			words = removeWord(words, i)
			i--

		case word == "(bin)":
			value, _ := strconv.ParseInt(words[i-1], 2, 64)
			words[i-1] = strconv.Itoa(int(value))
			words = removeWord(words, i)
			i--

		case word == "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words = removeWord(words, i)
			i--

		case word == "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words = removeWord(words, i)
			i--

		case word == "(cap)":
			words[i-1] = capitalize(words[i-1])
			words = removeWord(words, i)
			i--

		// handle (cap, n), (low, n)
		case strings.HasPrefix(words[i], "(up,") || strings.HasPrefix(word, "(low,") || strings.HasPrefix(word, "(cap,"):
			// Handle (tag, n)
			numStr := strings.TrimRight(words[i+1], ")")
			n, _ := strconv.Atoi(numStr)
			tag := word[1:4]
			for j := 1; j <= n; j++ {
				idx := i - j
				if idx >= 0 {
					if tag == "up," {
						words[idx] = strings.ToUpper(words[idx])
					}
					if tag == "low" {
						words[idx] = strings.ToLower(words[idx])
					}
					if tag == "cap" {
						words[idx] = capitalize(words[idx])
					}
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}

	}

	result := strings.Join(words, " ")

	result = fixPunctuation(result)
	result = fixQuotes(result)

	// fix a -> an
	endword := strings.Fields(result)
	for i := 0; i < len(endword)-1; i++ {

		if strings.ToLower(endword[i]) == "a" {

			first := strings.ToLower(string(endword[i+1][0]))

			if strings.ContainsAny(first, "aeiouh") {
				if endword[i] == "A" {
					endword[i] = "An"
				} else {
					endword[i] = "an"
				}
			}
		}
	}

	return strings.Join(endword, " ")
}
