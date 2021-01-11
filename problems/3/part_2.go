package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	slopeArgs := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for _, v := range slopeArgs {
		fmt.Printf(
			"Number of trees encountered for slope Right %d, Down %d: %d\n", v[0], v[1],
			slopeTreeCounter(v[0], v[1]),
		)
	}
}

func slopeTreeCounter(xJump int, yJump int) int {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Unable to open input file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var treeCount int
	position := -xJump

	for scanner.Scan() {
		line := scanner.Text()
		position = (position + xJump) % len(line)

		if string(line[position]) == "#" {
			treeCount++
		}

		for i := 1; i < yJump; i++ {
			scanner.Scan()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return treeCount
}
