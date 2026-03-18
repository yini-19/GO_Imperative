// // func capitalizeWord(wordstr string) string {

// // 	runes := []rune(wordstr)
// // 	runes[0] = unicode.ToUpper(runes[0])

// // 	for i := 1; i < len(runes); i++ {
// // 		runes[i] = unicode.ToLower(runes[i])
// // 	}
// // 	return string(runes)
// // }

// // func capUpLow(text string) string {
// // 	words := strings.Fields(text)

// // 	for i, _ := range words {
// // 		if strings.HasPrefix(words[i], "(") {
// // 			marker := strings.Trim(words[i], "()")
// // 			parts := strings.Split(marker, ",")

// // 			if len(parts) == 2 {

// // 				command := strings.TrimSpace(parts[0])
// // 				numstr := strings.TrimSpace(parts[1])

// // 				var n int
// // 				fmt.Sscanf(numstr, "%d", &n)

// // 				for j := i - 1; j >= 0 && j >= i-n; j-- {
// // 					switch command {
// // 					case "up, n":
// // 						words[j] = strings.ToUpper(words[j])
// // 					case "low, n":
// // 						words[j] = strings.ToLower(words[j])
// // 					case "cap, n":
// // 						words[j] = capitalizeWord(words[j])

// // 					}

// // 				}
// // 			} else {
// // 				command := strings.TrimSpace(marker)

// // 				switch command {
// // 				case "up":
// // 					words[i-1] = strings.ToUpper(words[i-1])
// // 				case "low":
// // 					words[i-1] = strings.ToLower(words[i-1])
// // 				case "cap":
// // 					words[i-1] = capitalizeWord(words[i-1])

// // 				}

// // 			}
// // 			words[i] = ""

// // 		}

// // 	}
// // 	var result []string

// // 	for _, w := range words {
// // 		if w != "" {
// // 			result = append(result, w)
// // 		}
// // 	}
// // 	return strings.Join(result, " ")

// // }

package main

import (
	"strings"
)

func capitalise(words []string, n int) []string {
	if n > len(words) {
		n = len(words)
	}

	for i := len(words) - n; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return words
}

//if n == "," || n == "!" {
//	return true
//} else {
//	return false
//}
//}
