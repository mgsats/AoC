package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Unable to open input file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(emptyLineSplitFunc)
	var countsTotal int
	for scanner.Scan() {
		line := scanner.Text()
		count := make(map[rune]struct{})
		for _, c := range line {
			if _, ok := count[c]; !ok && unicode.IsLower(c) {
				count[c] = struct{}{}
			}
		}
		countsTotal += len(count)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of valid passports: %d\n", countsTotal)
}

// https://stackoverflow.com/questions/33068644/how-a-scanner-can-be-implemented-with-a-custom-split
func emptyLineSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the index of the input of two newlines.
	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), data, nil
	}

	return
}
