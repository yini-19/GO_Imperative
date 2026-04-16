package main

import (
	"strconv"
	"strings"
)

func ProcessText(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		switch words[i] {

		case "(hex)":
			value, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = strconv.Itoa(int(value))
			words = removeWord(words, i)
			i--

		case "(bin)":
			value, _ := strconv.ParseInt(words[i-1], 2, 64)
			words[i-1] = strconv.Itoa(int(value))
			words = removeWord(words, i)
			i--

		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words = removeWord(words, i)
			i--

		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words = removeWord(words, i)
			i--

		case "(cap)":
			words[i-1] = capitalize(words[i-1])
			words = removeWord(words, i)
			i--
		}

		// handle (up,2) (low,3) etc
		if strings.HasPrefix(words[i], "(up,") {
			n := extractNumber(words[i])
			for j := 1; j <= n; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = removeWord(words, i)
			i--
		}

		if strings.HasPrefix(words[i], "(low,") {
			n := extractNumber(words[i])
			for j := 1; j <= n; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = removeWord(words, i)
			i--
		}

		if strings.HasPrefix(words[i], "(cap,") {
			n := extractNumber(words[i])
			for j := 1; j <= n; j++ {
				words[i-j] = capitalize(words[i-j])
			}
			words = removeWord(words, i)
			i--
		}
	}

	// fix a -> an
	for i := 0; i < len(words)-1; i++ {

		if strings.ToLower(words[i]) == "a" {

			first := strings.ToLower(string(words[i+1][0]))

			if strings.ContainsAny(first, "aeiouh") {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	result := strings.Join(words, " ")

	result = fixPunctuation(result)
	result = fixQuotes(result)

	return result
}