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
	var highestID int
	for scanner.Scan() {
		line := scanner.Text()

		var rowStr, columnStr string
		if len(line) == 10 {
			rowStr = line[0:7]
			columnStr = line[7:10]
		} else {
			break
		}

		row, _ := strconv.ParseInt(
			strings.ReplaceAll(
				strings.ReplaceAll(rowStr, "F", "0"), "B", "1",
			), 2, 0)
		column, _ := strconv.ParseInt(
			strings.ReplaceAll(
				strings.ReplaceAll(columnStr, "L", "0"), "R", "1",
			), 2, 0)
		ID := row*8 + column

		if int(ID) > highestID {
			highestID = int(ID)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Highest ID: %d\n", highestID)

}
