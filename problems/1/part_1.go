package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Unable to open input file")
		os.Exit(1)
	}
	defer file.Close()

	var currNumber int
	prevNumbers := make(map[int]struct{})
	for {

		_, err := fmt.Fscanf(file, "%d\n", &currNumber)

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				break
			}
			os.Exit(1)
		}

		lookingFor := 2020 - currNumber
		if _, ok := prevNumbers[lookingFor]; ok {
			fmt.Printf("Solution: %d x %d = %d\n", currNumber, lookingFor, currNumber*lookingFor)
			return
		}
		prevNumbers[currNumber] = struct{}{}

	}

	fmt.Println("No matching pair of numbers found that add to 2020")
}
