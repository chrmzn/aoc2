package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// It's listed as a left/right splitter but it's really returning first and second...
func splitter(str string, sep string) (left, right string) {
	stringSplit := strings.Split(str, sep)
	left, right = stringSplit[0], stringSplit[1]
	return left, right
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	var validPasswords int

	for _, line := range fileTextLines {
		// fmt.Println(line)
		matcher, password := splitter(line, ": ")
		fmt.Printf("Match: '%s'; Password: '%s'\n", matcher, password)

		valueCounts, valueMatchStr := splitter(matcher, " ")
		fmt.Printf("Counts: '%s'; Matcher: '%s'\n", valueCounts, valueMatchStr)

		minCountStr, maxCountStr := splitter(valueCounts, "-")
		fmt.Printf("MinCount: '%s'; MaxCount: '%s'\n", minCountStr, maxCountStr)

		minCount, err := strconv.Atoi(minCountStr)
		if err != nil {
			panic(err)
		}
		maxCount, err := strconv.Atoi(maxCountStr)
		if err != nil {
			panic(err)
		}
		fmt.Printf("MinCount: %d; MaxCount: %d\n", minCount, maxCount)

		var counter int
		valueMatch := rune(valueMatchStr[0])

		for _, char := range password {
			if char == valueMatch {
				counter++
			}
		}

		fmt.Printf("counter: %d\n", counter)

		if minCount <= counter && counter <= maxCount {
			fmt.Println("Password is valid!")
			validPasswords++
		}

		fmt.Println("")
	}

	fmt.Printf("Total Valid Passwords: %d", validPasswords)
}
