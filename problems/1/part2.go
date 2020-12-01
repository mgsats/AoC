package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Unable to open input file")
		os.Exit(1)
	}
	defer file.Close()

	var currNumber int

	numbers := []int{}
	for {

		_, err := fmt.Fscanf(file, "%d\n", &currNumber)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		numbers = append(numbers, currNumber)

	}

	sort.Ints(numbers)

	for i, n := range numbers {
		targetSum := 2020 - n
		j := i + 1
		k := len(numbers) - 1
		for j < k {
			currSum := numbers[j] + numbers[k]
			if currSum < targetSum {
				j += 1
			} else if currSum > targetSum {
				k -= 1
			} else {
				fmt.Printf("Solution:  %d x %d x %d = %d\n",
					n,
					numbers[j],
					numbers[k],
					n*numbers[j]*numbers[k],
				)
				return
			}
		}
	}

	fmt.Println("No matching pair of numbers found that add to 2020")
}
