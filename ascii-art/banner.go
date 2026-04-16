package main

import (
	"bufio"
	"os"
)

func LoadBanner(bannerFile string) (map[rune][]string, error) {
	filecont, err := os.Open(bannerFile)
	if err != nil {
		return nil, err
	}
	defer filecont.Close()

	bannerMap := make(map[rune][]string)
	scanner := bufio.NewScanner(filecont)

	var lines []string
	charcode := 32

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && len(lines) > 0 {

		}
	}

}
