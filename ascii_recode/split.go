package main

import "strings"

func SplitLine(line string) []string {
	return strings.Split(line, "\\n")
}