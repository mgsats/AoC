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
	prevNumbers := []int{}
	sums := make(map[int][2]int)
	for {

		_, err := fmt.Fscanf(file, "%d\n", &currNumber)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		lookingFor := 2020 - currNumber
		if otherNums, ok := sums[lookingFor]; ok {
			fmt.Printf("Solution:  %d x %d x %d = %d\n",
				currNumber,
				otherNums[0],
				otherNums[1],
				currNumber*otherNums[0]*otherNums[1])
			return
		}

		for _, n := range prevNumbers {
			sums[n+currNumber] = [2]int{n, currNumber}
		}
		prevNumbers = append(prevNumbers, currNumber)

	}

	fmt.Println("No matching pair of numbers found that add to 2020")
}
