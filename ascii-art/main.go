package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . <string>", "or", "go run . <string> <banner_file>")
		return
	}

	bannerFile := "standard.txt"
	if len(os.Args) == 3 {
		bannerFile = os.Args[2]
	}
	inputstr := strings.ReplaceAll(os.Args[1], "\\n", "\n")

	if inputstr == "" {
		fmt.Println()
		return
	}
	if inputstr == "\n" {
		fmt.Println()
		return
	}
	bannerMap, err := LoadBanner("banner/" + bannerFile)
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}

	PrintAscii(inputstr, bannerMap)
}
