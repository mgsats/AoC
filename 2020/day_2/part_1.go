package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Unable to open input file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validSum := 0
	for scanner.Scan() {
		lineFields := strings.Fields(scanner.Text())

		charLimits := strings.Split(lineFields[0], "-")
		lowerCharLimit, _ := strconv.Atoi(charLimits[0])
		upperCharLimit, _ := strconv.Atoi(charLimits[1])

		charOfInterest := strings.TrimSuffix(lineFields[1], ":")
		password := lineFields[2]
		charCount := strings.Count(password, charOfInterest)

		if charCount <= upperCharLimit && charCount >= lowerCharLimit {
			validSum++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of valid passwords: %d", validSum)

}
