package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func splitter(str string, sep string) (left, right string) {

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
		stringSplit := strings.Split(line, ": ")
		match, password := stringSplit[0], stringSplit[1]
		fmt.Printf("Match: '%s'; Password: '%s'\n", match, password)

		stringSplit = strings.Split(match, " ")
		valueCount, valueMatch := stringSplit[0], stringSplit[1]
		fmt.Printf("Counts: '%s'; Matcher: '%s'\n", valueCount, valueMatch)

		fmt.Println("")
	}

}
