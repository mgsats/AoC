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

		charPositions := strings.Split(lineFields[0], "-")
		lowerCharPosition, _ := strconv.Atoi(charPositions[0])
		upperCharPosition, _ := strconv.Atoi(charPositions[1])

		charOfInterest := strings.TrimSuffix(lineFields[1], ":")
		password := lineFields[2]

		firstPositionCheck := lowerCharPosition <= len(password) && string(password[lowerCharPosition-1]) == charOfInterest
		secondPositionCheck := upperCharPosition <= len(password) && string(password[upperCharPosition-1]) == charOfInterest
		if firstPositionCheck != secondPositionCheck {
			validSum++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of valid passwords: %d", validSum)

}
