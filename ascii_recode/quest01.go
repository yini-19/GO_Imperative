package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lineCount int
	var charCount int
	var wordCount int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lineCount++
		charCount += len(line)
		wordCount += strings.Count(line, " ") + 1
	}
	fmt.Println()
	fmt.Println("No of Lines:", lineCount)
	fmt.Println("No of Characters:", charCount)
	fmt.Println("No of Words:", wordCount)

}
