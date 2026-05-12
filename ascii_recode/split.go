package main

import "strings"

func SplitLine(line string) []string {
	line = strings.ReplaceAll(line, "\\n", "\n")
	return strings.Split(line, "\n")
}