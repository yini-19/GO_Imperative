package main

import (
	"bufio"
	"os"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bannerMap := make(map[rune][]string)
	scanner := bufio.NewScanner(file)

	var lines []string
	charcode := 32

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && len(lines) > 0 {
			bannerMap[rune(charcode)] = lines
			lines = []string{}
			charcode++
		}
		lines = append(lines, line)
	}
	return bannerMap, nil
}
