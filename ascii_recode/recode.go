package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <banner_file> <input_string>")
		return
	}

	bannerFile := os.Args[1]
	inputString := strings.ReplaceAll(os.Args[2], "\\n", "\n")

	bannerMap, err := LoadBanner("banner/" + bannerFile + ".txt")
	if err != nil {
		fmt.Printf("Error loading banner: %v\n", err)
		return
	}

	PrintAscii(inputString, bannerMap)
}
