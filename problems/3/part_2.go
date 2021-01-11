package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	slopeFile, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	slope := strings.Split(string(slopeFile), "\n")

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
			slopeTreeCounter(slope, v[0], v[1]),
		)
	}
}

func slopeTreeCounter(slope []string, xJump int, yJump int) int {
	var treeCount, positionX, positionY int

	for positionY < len(slope) && positionX < len(slope[positionY]) {
		if string(slope[positionY][positionX]) == "#" {
			treeCount++
		}
		positionX = (positionX + xJump) % len(slope[positionY])
		positionY += yJump
	}

	return treeCount
}
