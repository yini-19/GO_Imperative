package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func blockReader(filename string) map[string][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	blockMap := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	var block []string
	block1 := 1
	for scanner.Scan() {
		line := scanner.Text()
		block = append(block, line+"\n")
		if line == "" && block != nil {
			blockMap[fmt.Sprintf("block%d", block1)] = block
			for _, line := range block {
				fmt.Print(line)
			}
			block = []string{}
			block1++
		}
		
	}
	// fmt.Println(blockMap["block1"])
	// //fmt.Println()
	// fmt.Println(blockMap["block2"])
	// //fmt.Println()
	return blockMap
}

func main() {
	blockReader("example.txt")
	ReadFile("example.txt")
	

	// for line := 0; line < 5; line++ {
	// 	fmt.Println(blockMap[block1][line])
	// }
	// fmt.Println(blockReader("example.txt"))
	
}
