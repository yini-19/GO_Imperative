package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error! Usage: go run . <bannerfile> <string>")
		return
	}

	bannerFile := os.Args[1]
	inputstr := strings.ReplaceAll(os.Args[2], "\\n", "\n")

	bannerMap, err := LoadBanner("banner/" + bannerFile + ".txt")
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}

	PrintAscii(inputstr, bannerMap)
}
