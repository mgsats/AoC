package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Unable to open input file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var treeCount int
	position := -3
	for scanner.Scan() {
		line := scanner.Text()
		position = (position + 3) % len(line)
		if string(line[position]) == "#" {
			treeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of trees encountered: %d", treeCount)

}
