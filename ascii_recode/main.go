package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNERFILE]")
		return
	}

	option := os.Args[1]

	if !strings.HasPrefix(option, "--output=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNERFILE]")
		return
	}

	text := os.Args[2]
	bannerFile := os.Args[3]

	filename := strings.TrimPrefix(option, "--output=")
	fmt.Printf("Output file: %s\n", filename)

	banner, err := LoadBanner(bannerFile)
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}

	generated := GenerateArt(text, banner)
	err = os.WriteFile(filename, []byte(generated), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully!")
}
