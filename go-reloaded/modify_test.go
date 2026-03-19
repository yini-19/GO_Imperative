package main

import (
	"testing"
)

func TestProcessText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Hexadecimal Conversion",
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
		{
			name:     "Binary Conversion",
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},
		{
			name:     "Single Word Upper",
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO!",
		},
		{
			name:     "Single Word Lower",
			input:    "I should stop SHOUTING (low)",
			expected: "I should stop shouting",
		},
		{
			name:     "Single Word Capitalization",
			input:    "Welcome to the Brooklyn bridge (cap)",
			expected: "Welcome to the Brooklyn Bridge",
		},
		{
			name:     "Multi-word Uppercase",
			input:    "This is so exciting (up, 2)",
			expected: "This is SO EXCITING",
		},
		{
			name:     "Complex Multi-word Cap",
			input:    "it was the age of foolishness (cap, 6)",
			expected: "It Was The Age Of Foolishness",
		},
		{
			name:     "Standard Punctuation Spacing",
			input:    "I was sitting over there ,and then BAMM !!",
			expected: "I was sitting over there, and then BAMM!!",
		},
		{
			name:     "Ellipsis and Multiple Punctuation",
			input:    "I was thinking ... You were right",
			expected: "I was thinking... You were right",
		},
		{
			name:     "Single Word Quotes",
			input:    "I am exactly how they describe me: ' awesome '",
			expected: "I am exactly how they describe me: 'awesome'",
		},
		{
			name:     "Multi-word Quotes",
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			name:     "A to An before Vowel",
			input:    "There it was. A amazing rock!",
			expected: "There it was. An amazing rock!",
		},
		{
			name:     "A to An before H",
			input:    "There is no greater agony than bearing a untold story inside you.",
			expected: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:     "Mixed Requirements",
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ProcessText(tt.input)
			if actual != tt.expected {
				t.Errorf("\nFAILED: %s\nEXPECTED: %s\nACTUAL:   %s", tt.name, tt.expected, actual)
			}
		})
	}
}
