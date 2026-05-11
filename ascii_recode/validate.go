package main

import "fmt"

func ValidateInput(input string) (rune, error) {
	for _, char := range input {
		if char < 32 || char > 126 {
			return 0, fmt.Errorf("invalid character: %c", char)
		}
	}
	return 0, nil
}
