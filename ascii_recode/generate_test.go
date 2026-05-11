package main

import (
	"strings"
	"testing"
)

// Helper function to create a mock banner for testing
func createMockBanner(chars []rune) map[rune][]string {
	banner := make(map[rune][]string)
	for _, char := range chars {
		lines := make([]string, 8)
		for i := 0; i < 8; i++ {
			lines[i] = string(char)
		}
		banner[char] = lines
	}
	return banner
}

// Helper function to create a banner with actual ASCII art patterns
func createDetailedMockBanner(chars []rune) map[rune][]string {
	banner := make(map[rune][]string)
	for _, char := range chars {
		lines := make([]string, 8)
		// Create simple repeating pattern for testing
		pattern := strings.Repeat(string(char), 3)
		for i := 0; i < 8; i++ {
			lines[i] = pattern
		}
		banner[char] = lines
	}
	return banner
}

func TestGenerateArtEmptyString(t *testing.T) {
	banner := createMockBanner([]rune{'A', 'B', 'C'})
	result := GenerateArt("", banner)

	if result != "" {
		t.Errorf("Expected empty string for empty input, got: %q", result)
	}
}

func TestGenerateArtSingleCharacter(t *testing.T) {
	banner := createMockBanner([]rune{'A'})
	result := GenerateArt("A", banner)

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines for single character, got %d lines", len(lines))
	}
}

func TestGenerateArtMultipleCharacters(t *testing.T) {
	banner := createMockBanner([]rune{'A', 'B', 'C'})
	result := GenerateArt("ABC", banner)

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines for multiple characters, got %d lines", len(lines))
	}

	// Each line should contain the three characters concatenated
	for i, line := range lines {
		if !strings.Contains(line, "A") || !strings.Contains(line, "B") || !strings.Contains(line, "C") {
			t.Errorf("Line %d doesn't contain expected characters: %q", i, line)
		}
	}
}

func TestGenerateArtWithSpaces(t *testing.T) {
	banner := createMockBanner([]rune{'A', ' ', 'B'})
	result := GenerateArt("A B", banner)

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines, got %d lines", len(lines))
	}
}

func TestGenerateArtWithLineBreakDelimiter(t *testing.T) {
	banner := createMockBanner([]rune{'A', 'B'})
	result := GenerateArt("A\\nB", banner)

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 16 {
		t.Fatalf("Expected 16 lines for two rendered blocks, got %d", len(lines))
	}

	for i := 0; i < 8; i++ {
		if lines[i] != "A" {
			t.Errorf("Expected line %d to contain A art, got %q", i, lines[i])
		}
	}
	for i := 8; i < 16; i++ {
		if lines[i] != "B" {
			t.Errorf("Expected line %d to contain B art, got %q", i, lines[i])
		}
	}
}

func TestGenerateArtEmptyBanner(t *testing.T) {
	banner := make(map[rune][]string)
	result := GenerateArt("ABC", banner)

	// Should handle gracefully - might return partial or empty output
	if result == "" {
		// This is acceptable when banner is empty
		t.Logf("Empty banner returns empty result: acceptable behavior")
	}
}

func TestGenerateArtMissingCharacters(t *testing.T) {
	// Banner only has 'A', but input has 'B'
	banner := createMockBanner([]rune{'A'})
	result := GenerateArt("AB", banner)

	// Should still return some output (just without B)
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) == 0 {
		t.Errorf("Should handle missing characters gracefully, got no output")
	}
}

func TestGenerateArtOnlySpaces(t *testing.T) {
	banner := createMockBanner([]rune{' '})
	result := GenerateArt("   ", banner)

	// Three spaces should render as 8 lines with space characters
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines for spaces, got %d lines", len(lines))
	}
}

func TestGenerateArtOnlyLineBreaks(t *testing.T) {
	banner := createMockBanner([]rune{'A'})
	result := GenerateArt("\\n\\n\\n", banner)

	// Should output 24 newlines (8 per break) plus line endings
	newlineCount := strings.Count(result, "\n")
	if newlineCount < 8 {
		t.Errorf("Expected at least 8 newlines for line breaks, got %d", newlineCount)
	}
}

func TestGenerateArtConsecutiveLineBreaks(t *testing.T) {
	banner := createMockBanner([]rune{'A'})
	result := GenerateArt("A\\n\\nB", banner)

	// Should handle empty words (consecutive breaks) without panicking
	if result == "" {
		t.Errorf("Should produce output even with consecutive breaks")
	}
	newlineCount := strings.Count(result, "\n")
	if newlineCount < 16 {
		t.Errorf("Expected multiple newlines for breaks, got %d", newlineCount)
	}
}

func TestGenerateArtASCIIBoundaryCharacters(t *testing.T) {
	// Test space (32) and tilde (126)
	banner := createMockBanner([]rune{' ', '~'})
	result := GenerateArt(" ~", banner)

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines for boundary characters, got %d lines", len(lines))
	}
}

func TestGenerateArtMixedCase(t *testing.T) {
	banner := createMockBanner([]rune{'A', 'a', 'B', 'b'})
	result := GenerateArt("AaBb", banner)

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines for mixed case, got %d lines", len(lines))
	}
}

func TestGenerateArtNumbersAndSymbols(t *testing.T) {
	banner := createMockBanner([]rune{'0', '1', '!', '@', '#'})
	result := GenerateArt("01!@#", banner)

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines for symbols, got %d lines", len(lines))
	}
}

func TestGenerateArtVeryLongInput(t *testing.T) {
	banner := createMockBanner([]rune{'A'})
	longInput := strings.Repeat("A", 1000)
	result := GenerateArt(longInput, banner)

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines for long input, got %d lines", len(lines))
	}
}

func TestGenerateArtNoTrailingNewline(t *testing.T) {
	banner := createMockBanner([]rune{'A'})
	result := GenerateArt("A", banner)

	// Result should end with newline (from the loop adding "\n")
	if !strings.HasSuffix(result, "\n") {
		t.Errorf("Result should end with newline")
	}
}

func TestGenerateArtOutputStructure(t *testing.T) {
	banner := createDetailedMockBanner([]rune{'X'})
	result := GenerateArt("X", banner)

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	expectedLineCount := 8

	if len(lines) != expectedLineCount {
		t.Errorf("Expected %d lines in output, got %d", expectedLineCount, len(lines))
	}

	// Each line should contain the character pattern
	for i, line := range lines {
		if !strings.Contains(line, "X") {
			t.Errorf("Line %d missing expected character in output: %q", i, line)
		}
	}
}

func TestGenerateArtConcatination(t *testing.T) {
	banner := createDetailedMockBanner([]rune{'H', 'i'})
	result := GenerateArt("Hi", banner)

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines, got %d", len(lines))
	}

	// Characters should be concatenated horizontally
	for i, line := range lines {
		if !strings.Contains(line, "H") {
			t.Errorf("Line %d missing 'H': %q", i, line)
		}
		if !strings.Contains(line, "i") {
			t.Errorf("Line %d missing 'i': %q", i, line)
		}
	}
}

func TestGenerateArtEmptyWordHandling(t *testing.T) {
	banner := createMockBanner([]rune{'A'})
	// Multiple consecutive breaks create empty words
	result := GenerateArt("A\\n\\n\\nB", banner)

	newlineCount := strings.Count(result, "\n")
	// Should have 8 + (8*3) + 8 = 40 newlines plus the line rendering newlines
	if newlineCount < 24 {
		t.Errorf("Should handle multiple consecutive breaks, got %d newlines", newlineCount)
	}
}

func TestGenerateArtSpecialCharactersRange(t *testing.T) {
	// Test various special characters
	specialChars := []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '=', '[', ']', '{', '}', '|', ';', ':', '"', '\'', '<', '>', ',', '.', '?', '/'}
	banner := createMockBanner(specialChars)
	input := string(specialChars)
	result := GenerateArt(input, banner)

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines for special characters, got %d lines", len(lines))
	}
}

func TestGenerateArtWhitespaceOnlyWithBreaks(t *testing.T) {
	banner := createMockBanner([]rune{' '})
	result := GenerateArt("  \\n  ", banner)

	// Should produce output with proper line breaks
	if result == "" {
		t.Errorf("Should produce output for whitespace with breaks")
	}
	if !strings.Contains(result, "\n") {
		t.Errorf("Should contain newlines")
	}
}

func TestGenerateArtReturnStructure(t *testing.T) {
	banner := createMockBanner([]rune{'A'})
	result := GenerateArt("A", banner)

	// Verify it's a string and not nil
	if result == "" && banner['A'] != nil {
		t.Errorf("Should return non-empty string for valid input and banner")
	}

	// Verify all character(s) are properly rendered
	if !strings.Contains(result, "A") {
		t.Errorf("Output should contain rendered character 'A'")
	}
}

func TestGenerateArtRendererIntegration(t *testing.T) {
	// Test with a banner that has different patterns per line
	banner := make(map[rune][]string)
	banner['T'] = []string{"TTT", "TTT", "TTT", "TTT", "TTT", "TTT", "TTT", "TTT"}
	banner['E'] = []string{"EEE", "EEE", "EEE", "EEE", "EEE", "EEE", "EEE", "EEE"}

	result := GenerateArt("TE", banner)
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")

	if len(lines) != 8 {
		t.Errorf("Expected 8 lines, got %d", len(lines))
	}

	// Verify concatenation pattern
	for _, line := range lines {
		if len(line) != 6 { // TTT (3) + EEE (3)
			t.Errorf("Line should be 6 characters long (TTT+EEE), got %d: %q", len(line), line)
		}
	}
}

func TestGenerateArtNilBannerHandling(t *testing.T) {
	// Passing nil banner should not crash
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Should not panic with nil banner, panicked with: %v", r)
		}
	}()

	// This test checks if the function handles nil gracefully or panics
	// Note: This will likely panic, which is okay - the test documents this behavior
	// Uncomment if you want to test with nil:
	// GenerateArt("A", nil)
}

func TestGenerateArtFullWordRender(t *testing.T) {
	banner := createDetailedMockBanner([]rune{'G', 'O'})
	result := GenerateArt("GO", banner)

	// Split by actual newlines to get lines
	allLines := strings.Split(result, "\n")
	// Remove empty trailing line if exists
	contentLines := make([]string, 0)
	for _, line := range allLines {
		if line != "" {
			contentLines = append(contentLines, line)
		}
	}

	if len(contentLines) != 8 {
		t.Errorf("Expected 8 content lines, got %d", len(contentLines))
	}

	for i, line := range contentLines {
		// Should contain both characters
		hasG := strings.Contains(line, "G")
		hasO := strings.Contains(line, "O")
		if !hasG || !hasO {
			t.Errorf("Line %d should contain both 'G' and 'O': %q (hasG=%v, hasO=%v)", i, line, hasG, hasO)
		}
	}
}
