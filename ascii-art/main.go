package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error! Usage: go run main.go <bannerfile> <string>")
		return
	}

	bannerFile := os.Args[1]
	inputstr := os.Args[2]

	bannerMap, err := LoadBanner("banner/" + bannerFile + ".txt")
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}

	PrintAscii(inputstr, bannerMap)
}
