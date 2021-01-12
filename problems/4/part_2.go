package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	fieldRules := map[string]func(string) bool{
		"byr": func(str string) bool {
			year, err := strconv.Atoi(str)
			return len(str) == 4 &&
				err == nil &&
				year >= 1920 &&
				year <= 2002
		},
		"iyr": func(str string) bool {
			year, err := strconv.Atoi(str)
			return len(str) == 4 &&
				err == nil &&
				year >= 2010 &&
				year <= 2020
		},
		"eyr": func(str string) bool {
			year, err := strconv.Atoi(str)
			return len(str) == 4 &&
				err == nil &&
				year >= 2020 &&
				year <= 2030
		},
		"hgt": func(str string) bool {
			num, err := strconv.Atoi(str[:len(str)-2])
			units := str[len(str)-2:]
			if err != nil {
				return false
			} else if string(units) == "cm" {
				return num >= 150 && num <= 193
			} else if string(units) == "in" {
				return num >= 59 && num <= 76
			} else {
				return false
			}
		},
		"hcl": func(str string) bool {
			validColor, err := regexp.MatchString(`^\#[a-z0-9]{6}$`, str)
			return err == nil && validColor
		},
		"ecl": func(str string) bool {
			validColor, err := regexp.MatchString(`^(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)$`, str)
			return err == nil && validColor
		},
		"pid": func(str string) bool {
			validColor, err := regexp.MatchString(`^[0-9]{9}$`, str)
			return err == nil && validColor
		},
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(emptyLineSplitFunc)
	var validPassportCount int
	for scanner.Scan() {
		line := scanner.Text()
		lineFields := strings.Fields(line)

		containsAllFields := true
		for s, _ := range fieldRules {
			if !strings.Contains(line, s) {
				containsAllFields = false
				break
			}
		}

		allFieldsValid := true
		for _, s := range lineFields {
			fieldAndValue := strings.Split(s, ":")
			if len(fieldAndValue) != 2 {
				allFieldsValid = false
				break
			}
			field, value := fieldAndValue[0], fieldAndValue[1]
			if fn, ok := fieldRules[field]; ok && !fn(value) {
				allFieldsValid = false
				break
			}
		}

		if containsAllFields && allFieldsValid {
			validPassportCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of valid passports: %d\n", validPassportCount)
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
