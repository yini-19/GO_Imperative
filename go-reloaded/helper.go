package main

import (
	"strconv"
	"strings"
)

func removeWord(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice // return unchanged if out of bounds
	}
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

	punctuations := []string{".", ",", "!", "?", ":", ";"}

	for _, p := range punctuations {
		text = strings.ReplaceAll(text, " "+p, p+" ")
		text = strings.ReplaceAll(text, p+" ", p+" ")
	}

	return text

}

// func fixPunctuation(tokens []string) string {
// 	var builder strings.Builder

// 	for i, token := range tokens {
// 		runes := []rune(token)

// 		if len(runes) == 1 && unicode.IsPunct(runes[0]) {
// 			builder.WriteString(token)

// 			if i+1 < len(tokens) {
// 				nextRunes := []rune(tokens[i+1])
// 				if !(len(nextRunes) == 1 && unicode.IsPunct(nextRunes[0])) {
// 					builder.WriteRune(' ')
// 				}
// 			}
// 		} else {

// 			if i > 0 {
// 				prevRunes := []rune(tokens[i-1])
// 				if !(len(prevRunes) == 1 && unicode.IsPunct(prevRunes[0])) {
// 					builder.WriteRune(' ')
// 				}
// 			}
// 			builder.WriteString(token)
// 		}
// 	}
// 	return builder.String()
// }

func fixQuotes(text string) string {

	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")

	// fix cases where punctuation comes after the quote
	text = strings.ReplaceAll(text, " '.", "'.")
	text = strings.ReplaceAll(text, " ',", "',")
	text = strings.ReplaceAll(text, " '!", "'!")
	text = strings.ReplaceAll(text, " '?", "'?")
	text = strings.ReplaceAll(text, ".'", ". '")
	text = strings.ReplaceAll(text, "' .", "'.")
	text = strings.ReplaceAll(text, ",'", ", '")
	text = strings.ReplaceAll(text, ":'", ": '")
	text = strings.ReplaceAll(text, ";'", "; '")
	text = strings.ReplaceAll(text, ". .", "..")
	text = strings.ReplaceAll(text, "! !", "!!")

	return text
}
