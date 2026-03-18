package main

import (
	"strconv"
	"strings"
)

func removeWord(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func capitalize(word string) string {

	if len(word) == 0 {
		return word
	}

	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func extractNumber(tag string) int {

	tag = strings.Trim(tag, "()")

	parts := strings.Split(tag, ",")

	n, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

	return n
}

func fixPunctuation(text string) string {

	text = strings.ReplaceAll(text, " ,", ",")
	text = strings.ReplaceAll(text, " .", ".")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, " :", ":")
	text = strings.ReplaceAll(text, " ;", ";")

	return text
}

func fixQuotes(text string) string {

	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")

	return text
}
