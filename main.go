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

	for _, line := range fileTextLines {
		// fmt.Println(line)
		matcher, password := splitter(line, ": ")
		fmt.Printf("Match: '%s'; Password: '%s'\n", matcher, password)

		valueCounts, valueMatch := splitter(matcher, " ")
		fmt.Printf("Counts: '%s'; Matcher: '%s'\n", valueCounts, valueMatch)

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

		fmt.Println("")
	}

}
